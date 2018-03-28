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

func mergeMap(m1, m2 map[string]string) map[string]string {
	result := make(map[string]string, len(m1)+len(m2))

	for key, val := range m1 {
		result[key] = val
	}

	for key, val := range m2 {
		if _, ok := result[key]; ok {
			continue
		}
		result[key] = val
	}
	return result
}
