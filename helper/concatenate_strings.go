package helper

import "strings"

// ConcatenateStrings takes a number of string arguments and concatenates them
func ConcatenateStrings(tokens ...string) string {
	var result strings.Builder
	for _, token := range tokens {
		result.WriteString(token)
	}
	return result.String()
}
