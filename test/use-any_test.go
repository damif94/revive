package test

import (
	"testing"

	"github.com/damif94/revive/rule"
)

func TestUseAny(t *testing.T) {
	testRule(t, "use-any", &rule.UseAnyRule{})
}
