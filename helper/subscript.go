package helper

import (
	"regexp"
)

func Subscript(s string) string {
	re := regexp.MustCompile(`~(\S+)~`)
	in := []byte(s)
	out := re.ReplaceAllFunc(in, func(s []byte) []byte {
		match := re.ReplaceAllString(string(s), `$1`)
		return []byte("<sub>" + match + "</sub>")
	})
	return string(out)
}
