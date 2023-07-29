package test

import (
	"testing"

	"github.com/damif94/revive/lint"
	"github.com/damif94/revive/rule"
)

func TestCommentSpacings(t *testing.T) {

	testRule(t, "comment-spacings", &rule.CommentSpacingsRule{}, &lint.RuleConfig{
		Arguments: []interface{}{"myOwnDirective"}},
	)
}
