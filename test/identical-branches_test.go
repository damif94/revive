package test

import (
	"testing"

	"github.com/damif94/revive/rule"
)

// IdenticalBranches rule.
func TestIdenticalBranches(t *testing.T) {
	testRule(t, "identical-branches", &rule.IdenticalBranchesRule{})
}
