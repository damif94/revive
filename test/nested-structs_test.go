package test

import (
	"testing"

	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

func TestNestedStructs(t *testing.T) {
	testRule(t, "nested-structs", &rule.NestedStructs{}, &lint.RuleConfig{})
}
