package test

import (
	"testing"

	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

func TestErrorStringsWithCustomFunctions(t *testing.T) {
	args := []interface{}{"pkgErrors.Wrap"}
	testRule(t, "error-strings-with-custom-functions", &rule.ErrorStringsRule{}, &lint.RuleConfig{
		Arguments: args,
	})
}
