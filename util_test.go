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
		got := getElementCountOfMaps(test.m1, test.m2)

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
		got := mergeMap(test.m1, test.m2)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}

func TestCreateOneHotVectorFromText(t *testing.T) {
	tests := []struct {
		text     string
		expected map[string]string
	}{
		{
			text: "お寿司が食べたい",
			expected: map[string]string{
				"お":  "10000",
				"寿司": "01000",
				"が":  "00100",
				"食べ": "00010",
				"たい": "00001",
			},
		},
		{
			text: "私はノートを買ったが、買い間違えた",
			expected: map[string]string{
				"私":   "1000000000",
				"は":   "0100000000",
				"ノート": "0010000000",
				"を":   "0001000000",
				"買っ":  "0000100000",
				"た":   "0000010000",
				"が":   "0000001000",
				"、":   "0000000100",
				"買い":  "0000000010",
				"間違え": "0000000001",
			},
		},
	}

	for _, test := range tests {
		got := CreateOneHotVectorFromText(test.text)
		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("expected: %v, got: %v", test.expected, got)
		}
	}
}
