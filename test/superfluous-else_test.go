package test

import (
	"testing"

	"github.com/damif94/revive/internal/ifelse"
	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

// TestSuperfluousElse rule.
func TestSuperfluousElse(t *testing.T) {
	testRule(t, "superfluous-else", &rule.SuperfluousElseRule{})
	testRule(t, "superfluous-else-scope", &rule.SuperfluousElseRule{}, &lint.RuleConfig{Arguments: []interface{}{ifelse.PreserveScope}})
}
