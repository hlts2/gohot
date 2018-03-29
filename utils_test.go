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
		got := getUniqTokens(test.tokens)

		if !reflect.DeepEqual(got, getUniqTokens(test.tokens)) {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}

func TestGetElementCountOfMaps(t *testing.T) {
	tests := []struct {
		m1       map[string]string
		m2       map[string]string
		expected int
	}{
		{
			m1:       map[string]string{"name": "hlts2", "age": "22"},
			m2:       map[string]string{"loc": "Osaka"},
			expected: 3,
		},
		{
			m1:       map[string]string{"name": "hlts2", "age": "22"},
			m2:       map[string]string{"name": "hiroto"},
			expected: 2,
		},
		{
			m1:       map[string]string{},
			m2:       map[string]string{},
			expected: 0,
		},
	}

	for _, test := range tests {
		got := getElementCount(test.m1, test.m2)

		if test.expected != got {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}

func TestMergeMap(t *testing.T) {
	tests := []struct {
		m1       map[string]string
		m2       map[string]string
		expected map[string]string
	}{
		{
			m1:       map[string]string{"name": "hlts2"},
			m2:       map[string]string{"age": "23"},
			expected: map[string]string{"name": "hlts2", "age": "23"},
		},
		{
			m1:       map[string]string{"name": "hlts2"},
			m2:       map[string]string{"name": "hiroto", "age": "23"},
			expected: map[string]string{"name": "hlts2", "age": "23"},
		},
		{
			m1:       map[string]string{},
			m2:       map[string]string{},
			expected: map[string]string{},
		},
	}

	for _, test := range tests {
		got := merge(test.m1, test.m2)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}
