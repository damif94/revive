package test

import (
	"testing"

	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

func TestAddConstant(t *testing.T) {
	args := []interface{}{map[string]interface{}{
		"maxLitCount": "2",
		"allowStrs":   "\"\"",
		"allowInts":   "0,1,2",
		"allowFloats": "0.0,1.0",
		"ignoreFuncs": "os\\.(CreateFile|WriteFile|Chmod|FindProcess),\\.Println,ignoredFunc",
	}}

	testRule(t, "add-constant", &rule.AddConstantRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
