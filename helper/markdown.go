package helper

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
)

// Markdown converts markdown to HTML
func Markdown(md string) template.HTML {
	// lot's of the markdown we currently have stored doesn't match markdown title specs
	// currently it has no space between the hashes and the title text e.g. ##Title
	// to use our new markdown parser we have add a space e.g. ## Title
	re := regexp.MustCompile(`(##+)([^\s#])`)

	modifiedMarkdown := strings.Builder{}
	for _, line := range strings.Split(md, "\n") {
		modifiedMarkdown.WriteString(fmt.Sprintf("%s\n", re.ReplaceAllString(line, "$1 $2")))
	}

	s := blackfriday.Run([]byte(modifiedMarkdown.String()))
	return template.HTML(s)
}
