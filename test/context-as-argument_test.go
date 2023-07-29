package test

import (
	"testing"

	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

func TestContextAsArgument(t *testing.T) {
	testRule(t, "context-as-argument", &rule.ContextAsArgumentRule{}, &lint.RuleConfig{
		Arguments: []interface{}{
			map[string]interface{}{
				"allowTypesBefore": "AllowedBeforeType,AllowedBeforeStruct,*AllowedBeforePtrStruct,*testing.T",
			},
		},
	},
	)
}
