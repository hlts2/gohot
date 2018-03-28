package gohot

func getUniqTokens(tokens []string) []string {
	if len(tokens) == 0 {
		return []string{}
	}

	uniqTokens := make([]string, len(tokens))
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

func getElementCount(m1, m2 map[string]string) int {
	duplicateCount := 0
	for m1Key := range m1 {
		for m2Key := range m2 {
			if m1Key == m2Key {
				duplicateCount++
			}
		}
	}
	return (len(m1) + len(m2)) - duplicateCount
}

func merge(m1, m2 map[string]string) map[string]string {
	result := make(map[string]string, getElementCount(m1, m2))

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
