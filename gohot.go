package gohot

// GetOneHotVectorOfTokens returns one hot vector of tokens
func GetOneHotVectorOfTokens(tokens []string) map[string]string {
	return getOneHotVectorOfTokens(tokens)
}

func getOneHotVectorOfTokens(tokens []string) map[string]string {
	uniqTokens := getRemovedDuplicateTokens(tokens)

	oneHotVectors := make(map[string]string, len(uniqTokens))

	for i, uniqToken := range uniqTokens {
		vector := ""

		for j := 0; j < len(uniqTokens); j++ {
			if j == i {
				vector += "1"
			} else {
				vector += "0"
			}
		}
		oneHotVectors[uniqToken] = vector
	}
	return oneHotVectors
}
