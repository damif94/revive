package test

import (
	"testing"

	"github.com/damif94/revive/internal/ifelse"
	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

// TestEarlyReturn tests early-return rule.
func TestEarlyReturn(t *testing.T) {
	testRule(t, "early-return", &rule.EarlyReturnRule{})
	testRule(t, "early-return-scope", &rule.EarlyReturnRule{}, &lint.RuleConfig{Arguments: []interface{}{ifelse.PreserveScope}})
}
