package gohot

import (
	"reflect"
	"testing"
)

func TestGetRemovedDuplicateTaokens(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected []string
	}{
		{
			tokens:   []string{"a", "b", "c", "d"},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			tokens:   []string{"a", "a"},
			expected: []string{"a"},
		},
		{
			tokens:   []string{},
			expected: []string{},
		},
	}

	for _, test := range tests {
		got := getRemovedDuplicateTokens(test.tokens)

		if !reflect.DeepEqual(got, getRemovedDuplicateTokens(test.tokens)) {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}
