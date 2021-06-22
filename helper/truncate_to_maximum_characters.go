package helper

import "strings"

// TruncateToMaximumCharacters returns a substring of parameter 'text' if the text is longer than the specified maximum length
func TruncateToMaximumCharacters(text string, maxLength int) string {
	if len(text) < maxLength {
		return text
	}

	truncatedText := text[0:maxLength]
	return strings.TrimSpace(truncatedText) + "..."
}
