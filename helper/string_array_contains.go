package helper

func StringArrayContains(comparitor string, items []string) bool {
	for _, token := range items {
		if token == comparitor {
			return true
		}
	}
	return false
}
