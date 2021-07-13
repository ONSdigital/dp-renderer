package helper

import "html/template"

func SafeHTML(s string) string {
	return template.HTMLEscapeString(s)
}
