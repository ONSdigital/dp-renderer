package helper_test

import (
	"html/template"
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMarkdown(t *testing.T) {
	Convey("markdown should return expected HTML", t, func() {
		for i := 0; i < len(testMarkdown); i++ {
			So(helper.Markdown(testMarkdown[i]), ShouldEqual, testHTMLForMarkdown[i])
		}
	})
}

var testMarkdown = []string{
	"- First bullet point\n\n- Second bullet point\n\n\tSecond line of second bullet point",
	"1. First bullet point\n\n2. Second bullet point",
	"##Heading2",
	"### Heading 3",
	"#### Heading 4 \n\n Some text, _some italic text_, **some strong text**",
	"[Test link](https://www.test.com)",
	"### Heading 3 \n\n some text [Anchor link](#)",
}

// the new line at the end of each test case is needed
var testHTMLForMarkdown = []template.HTML{
	"<ul>\n<li><p>First bullet point</p></li>\n\n<li><p>Second bullet point</p>\n\n<p>Second line of second bullet point</p></li>\n</ul>\n",
	"<ol>\n<li><p>First bullet point</p></li>\n\n<li><p>Second bullet point</p></li>\n</ol>\n",
	"<h2>Heading2</h2>\n",
	"<h3>Heading 3</h3>\n",
	"<h4>Heading 4</h4>\n\n<p>Some text, <em>some italic text</em>, <strong>some strong text</strong></p>\n",
	"<p><a href=\"https://www.test.com\">Test link</a></p>\n",
	"<h3>Heading 3</h3>\n\n<p>some text <a href=\"#\">Anchor link</a></p>\n",
}
