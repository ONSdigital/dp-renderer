package helper

import "strings"

// Join takes a slice of string arguments and concatenates them
func Join(sep string, s []string) string {
	return strings.Join(s, sep)
}
