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

	tokensObjects := t.Tokenize(text)
	tokens := make([]string, 0, len(tokensObjects))

	for _, tokensObject := range tokensObjects {
		if tokensObject.Class == tokenizer.DUMMY {
			continue
		}
		tokens = append(tokens, tokensObject.Surface)
	}

	uniqTokens := getUniqTokens(tokens)
	return createOneHotVectorFromTokens(uniqTokens)
}

func createOneHotVectorFromTokens(uniqTokens []string) map[string]string {
	oneHotVector := make(map[string]string, len(uniqTokens))

	for i, uniqToken := range uniqTokens {
		vector := make([]byte, 0, len(uniqTokens))

		for j := 0; j < len(uniqTokens); j++ {
			if j == i {
				vector = append(vector, "1"...)
				continue
			}
			vector = append(vector, "0"...)
		}
		oneHotVector[uniqToken] = string(vector)
	}
	return oneHotVector
}
