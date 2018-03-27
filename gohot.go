package gohot

// CreateOneHotVectorFromTokens returns one hot vector of tokens
func CreateOneHotVectorFromTokens(tokens []string) map[string]string {
	uniqTokens := getRemovedDuplicateTokens(tokens)

	return createOneHotVectorFromTokens(uniqTokens)
}

// CreateOneHotVectorFromText returns one hot vector of text
func CreateOneHotVectorFromText(text string) map[string]string {
	tokens := text2Tokens(text)

	uniqTokens := getRemovedDuplicateTokens(tokens)

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
