package test

import (
	"testing"

	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

func TestRangeValInClosure(t *testing.T) {
	testRule(t, "range-val-in-closure", &rule.RangeValInClosureRule{}, &lint.RuleConfig{})
}
