package gohot

func getRemovedDuplicateTokens(tokens []string) []string {
	if len(tokens) == 0 {
		return []string{}
	}

	uniqTokens := make([]string, len(tokens)+1)
	index := 0

	for _, token := range tokens {
		duplicate := false
		for _, uniqToken := range uniqTokens {
			if uniqToken == token {
				duplicate = true
				break
			}
		}

		if !duplicate {
			uniqTokens[index] = token
			index++
		}
	}
	return uniqTokens[:index]
}
