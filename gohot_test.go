package gohot

import (
	"reflect"
	"testing"
)

func TestCreateOneHotVectorOfTokens(t *testing.T) {
	tests := []struct {
		tokens   []string
		expected map[string]string
	}{
		{
			tokens: []string{"a", "b", "c"},
			expected: map[string]string{
				"a": "100",
				"b": "010",
				"c": "001",
			},
		},
		{
			tokens: []string{"a", "b", "a", "c"},
			expected: map[string]string{
				"a": "100",
				"b": "010",
				"c": "001",
			},
		},
		{
			tokens:   []string{},
			expected: map[string]string{},
		},
	}

	for _, test := range tests {
		oneHotVectors := CreateOneHotVectorFromTokens(test.tokens)

		if !reflect.DeepEqual(oneHotVectors, test.expected) {
			t.Errorf("expected: %v, got: %v", oneHotVectors, test.expected)
		}
	}
}
