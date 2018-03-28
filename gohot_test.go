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
