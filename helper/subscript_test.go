package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscript(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Expected    string
	}{
		{"'hello world' should become 'hello world'", "hello world", "hello world"},
		{"'~hello world~' should become '~hello world~'", "~hello world~", "~hello world~"},
		{"'~hello~ world' should become '<sub>hello</sub> world'", "~hello~ world", "<sub>hello</sub> world"},
		{"'hello ~world~' should become 'hello <sub>world</sub>'", "hello ~world~", "hello <sub>world</sub>"},
		{"'this ~is~ subscript' should become 'this <sub>is</sub> subscript'", "this ~is~ subscript", "this <sub>is</sub> subscript"},
		{"'~this~ is ~subscript~' should become '<sub>this</sub> is <sub>subscript</sub>'", "~this~ is ~subscript~", "<sub>this</sub> is <sub>subscript</sub>"},
	}

	for _, test := range cases {
		Convey(test.Description, t, func() {
			got := helper.Subscript(test.Input)
			So(got, ShouldEqual, test.Expected)
		})
	}
}
