package test

import (
	"testing"

	"github.com/damif94/revive/rule"
)

func TestUnconditionalRecursion(t *testing.T) {
	testRule(t, "unconditional-recursion", &rule.UnconditionalRecursionRule{})
}
