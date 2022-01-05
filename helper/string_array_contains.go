package helper

func StringArrayContains(s string, items []string) bool {
	for _, item := range items {
		if item == s {
			return true
		}
	}
	return false
}
