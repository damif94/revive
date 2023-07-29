package test

import (
	"testing"

	"github.com/damif94/revive/rule"
)

// UselessBreak rule.
func TestUselessBreak(t *testing.T) {
	testRule(t, "useless-break", &rule.UselessBreak{})
}
