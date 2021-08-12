package helper

import "strings"

// TrimPrefixedPeriod remove prefixed periods from the start of the string
func TrimPrefixedPeriod(input string) string {
	return strings.TrimLeft(input, ".")
}
