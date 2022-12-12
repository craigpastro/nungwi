package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRewriteParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		ok    bool
	}{
		{
			name:  "Self",
			input: "self",
			ok:    true,
		},
		{
			name:  "ComputedUserset1",
			input: "computedUserset(viewer)",
			ok:    true,
		},
		{
			name:  "ComputedUserset2",
			input: "computedUserset(parent, viewer)",
			ok:    false,
		},
		{
			name:  "TupleToUserset1",
			input: "tupleToUserset(parent, viewer)",
			ok:    true,
		},
		{
			name:  "TupleToUserset2",
			input: "tupleToUserset(computedUserset(parent), viewer)",
			ok:    false,
		},
		{
			name:  "Union1",
			input: "union(self, computedUserset(viewer))",
			ok:    true,
		},
		{
			name:  "Union2",
			input: "union(intersection(self, tupleToUserset(parent, viewer)), computedUserset(writer))",
			ok:    true,
		},
		{
			name:  "Intersection1",
			input: "intersection(computedUserset(viewer), self)",
			ok:    true,
		},
		{
			name:  "Intersection2",
			input: "union(computedUserset(viewer), computedUserset(writer), computedUserset(owner))",
			ok:    false,
		},
		{
			name:  "Exclusion1",
			input: "exclusion(computedUserset(viewer), computedUserset(banned))",
			ok:    true,
		},
		{
			name:  "Exclusion2",
			input: "union(tupleToUserset(parent, viewer), computedUserset(banned))",
			ok:    true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.ok, ParseRewrite(test.input))
		})
	}
}

func TestUserParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		ok    bool
	}{
		{
			name:  "UserID",
			input: "aardvark",
			ok:    true,
		},
		{
			name:  "Object",
			input: "object(folder, x)",
			ok:    true,
		},
		{
			name:  "Userset",
			input: "userset(group, eng, member)",
			ok:    true,
		},
		{
			name:  "BadUserId1",
			input: "foo(bar",
			ok:    false,
		},
		{
			name:  "BadUserId2",
			input: "foo*bar",
			ok:    false,
		},
		{
			name:  "BadObject",
			input: "object(groud, end, member)",
			ok:    false,
		},
		{
			name:  "BadUserset",
			input: "userset(folder, x)",
			ok:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.ok, ParseUser(test.input))
		})
	}
}
