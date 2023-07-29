package test

import (
	"testing"

	"github.com/damif94/revive/rule"
)

func TestUnreachableCode(t *testing.T) {
	testRule(t, "unreachable-code", &rule.UnreachableCodeRule{})
}
