package gconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchingRules(t *testing.T) {
	a := Role{Rules: []Rule{{Group: "developers"}}}

	type testcase struct {
		groups []string
		want   []Rule
	}

	testcases := []testcase{
		{groups: []string{"developers"}, want: []Rule{{Group: "developers"}}},
		{groups: []string{"other"}, want: nil},
	}

	for _, tc := range testcases {
		got := a.MatchingRules(tc.groups)
		assert.Equal(t, tc.want, got)
	}
}
