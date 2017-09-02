package linter

import (
	"go/ast"
	"go/token"
	"math"
	"regexp"
	"strings"

	"github.com/mgechev/revive/file"
	"github.com/mgechev/revive/rule"
)

// ReadFile defines an abstraction for reading files.
type ReadFile func(path string) (result []byte, err error)

// Linter is used for linting set of files.
type Linter struct {
	reader ReadFile
}

// New creates a new Linter
func New(reader ReadFile) Linter {
	return Linter{reader: reader}
}

// Lint lints a set of files with the specified rule.
func (l *Linter) Lint(filenames []string, ruleSet []rule.Rule) ([]rule.Failure, error) {
	var fileSet token.FileSet
	var failures []rule.Failure
	var ruleNames = []string{}
	for _, r := range ruleSet {
		ruleNames = append(ruleNames, r.GetName())
	}
	for _, filename := range filenames {
		content, err := l.reader(filename)
		if err != nil {
			return nil, err
		}
		file, err := file.New(filename, content, &fileSet)
		disabledIntervals := l.disabledIntervals(file, ruleNames)

		if err != nil {
			return nil, err
		}

		for _, rule := range ruleSet {
			currentFailures := rule.Apply(file, []string{})
			currentFailures = l.filterFailures(currentFailures, disabledIntervals)
			failures = append(failures, currentFailures...)
		}
	}

	return failures, nil
}

type enableDisableConfig struct {
	enabled  bool
	position int
}

func (l *Linter) disabledIntervals(file *file.File, allRuleNames []string) []rule.DisabledInterval {
	re := regexp.MustCompile(`^\s*revive:(enable|disable)(?:-(line|next-line))?(:|\s|$)`)

	enabledDisabledRulesMap := make(map[string][]enableDisableConfig)

	getEnabledDisabledIntervals := func() []rule.DisabledInterval {
		var result []rule.DisabledInterval

		for ruleName, disabledArr := range enabledDisabledRulesMap {
			for i := 0; i < len(disabledArr); i++ {
				interval := rule.DisabledInterval{
					RuleName: ruleName,
					From: token.Position{
						Filename: file.Name,
						Line:     disabledArr[i].position,
					},
					To: token.Position{
						Filename: file.Name,
						Line:     math.MaxInt32,
					},
				}
				if i%2 == 0 {
					result = append(result, interval)
				} else {
					result[len(result)-1].To.Line = disabledArr[i].position
				}
			}
		}
		return result
	}

	handleConfig := func(isEnabled bool, line int, name string) {
		existing, ok := enabledDisabledRulesMap[name]
		if !ok {
			existing = []enableDisableConfig{}
			enabledDisabledRulesMap[name] = existing
		}
		if (len(existing) > 1 && existing[len(existing)-1].enabled == isEnabled) ||
			(len(existing) == 0 && isEnabled) {
			return
		}
		existing = append(existing, enableDisableConfig{
			enabled:  isEnabled,
			position: line,
		})
		enabledDisabledRulesMap[name] = existing
	}

	handleRules := func(filename, modifier string, isEnabled bool, line int, ruleNames []string) []rule.DisabledInterval {
		var result []rule.DisabledInterval
		for _, name := range ruleNames {
			if modifier == "line" {
				handleConfig(!isEnabled, line, name)
				handleConfig(isEnabled, line, name)
			} else if modifier == "next-line" {
				handleConfig(!isEnabled, line+1, name)
				handleConfig(isEnabled, line+1, name)
			} else {
				handleConfig(isEnabled, line, name)
			}
		}
		return result
	}

	handleComment := func(filename string, c *ast.CommentGroup, line int) {
		text := c.Text()
		parts := re.FindStringSubmatch(text)
		if len(parts) == 0 {
			return
		}
		str := re.FindString(text)
		ruleNamesString := strings.Split(text, str)
		ruleNames := []string{}
		if len(ruleNamesString) == 2 {
			tempNames := strings.Split(ruleNamesString[1], ",")
			for _, name := range tempNames {
				if len(name) > 0 {
					ruleNames = append(ruleNames, name)
				}
			}
		}

		if len(ruleNames) == 0 {
			ruleNames = allRuleNames
		}

		handleRules(filename, parts[2], parts[1] == "enable", line, ruleNames)
	}

	comments := file.GetAST().Comments
	for _, c := range comments {
		handleComment(file.Name, c, file.ToPosition(c.Pos()).Line)
	}

	return getEnabledDisabledIntervals()
}

func (l *Linter) filterFailures(failures []rule.Failure, disabledIntervals []rule.DisabledInterval) []rule.Failure {
	return failures
}
