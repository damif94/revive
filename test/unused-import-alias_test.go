package test

import (
	"github.com/damif94/revive/rule"
	"testing"
)

// TestUnnecessaryStmt rule.
func TestUnusedImportAlias(t *testing.T) {
	testRule(t, "unused-import-alias", &rule.UnusedImportAlias{})
}
