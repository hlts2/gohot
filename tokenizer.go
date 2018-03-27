package gohot

import (
	"github.com/ikawaha/kagome/tokenizer"
)

func text2Tokens(text string) []string {
	t := tokenizer.New()

	tokens := []string{}
	for _, token := range t.Tokenize(text) {
		tokens = append(tokens, token.String())
	}
	return tokens
}
