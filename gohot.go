package gohot

import (
	"github.com/ikawaha/kagome/tokenizer"
)

// CreateOneHotVectorFromTokens returns one hot vector of tokens
func CreateOneHotVectorFromTokens(tokens []string) map[string]string {
	uniqTokens := getUniqTokens(tokens)
	return createOneHotVectorFromTokens(uniqTokens)
}

// CreateOneHotVectorFromText returns one hot vector of text
func CreateOneHotVectorFromText(text string) map[string]string {
	t := tokenizer.New()

	tokenObjects := t.Tokenize(text)
	tokens := make([]string, len(tokenObjects))

	for _, tokenObject := range tokenObjects {
		if tokenObject.Class == tokenizer.DUMMY {
			continue
		}
		tokens = append(tokens, tokenObject.Surface)
	}

	uniqTokens := getUniqTokens(tokens)

	return createOneHotVectorFromTokens(uniqTokens)
}

func createOneHotVectorFromTokens(uniqTokens []string) map[string]string {
	oneHotVector := make(map[string]string, len(uniqTokens))

	for i, uniqToken := range uniqTokens {
		vector := ""

		for j := 0; j < len(uniqTokens); j++ {
			if j == i {
				vector += "1"
			} else {
				vector += "0"
			}
		}
		oneHotVector[uniqToken] = vector
	}
	return oneHotVector
}
