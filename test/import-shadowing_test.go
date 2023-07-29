package test

import (
	"testing"

	"github.com/damif94/revive/rule"
)

func TestImportShadowing(t *testing.T) {
	testRule(t, "import-shadowing", &rule.ImportShadowingRule{})
}
