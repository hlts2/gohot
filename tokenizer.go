package gohot

import (
	"github.com/ikawaha/kagome/tokenizer"
)

func text2Tokens(text string) []string {
	t := tokenizer.New()

	tokens := t.Tokenize(text)
	result := make([]string, len(tokens))

	for _, token := range tokens {
		result = append(result, token.String())
	}
	return result
}
