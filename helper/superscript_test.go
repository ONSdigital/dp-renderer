package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSuperscript(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Expected    string
	}{
		{"'hello world' should become 'hello world'", "hello world", "hello world"},
		{"'^hello world^' should become '^hello world^'", "^hello world^", "^hello world^"},
		{"'^hello^ world' should become '<sup>hello</sup> world'", "^hello^ world", "<sup>hello</sup> world"},
		{"'hello ^world^' should become 'hello <sup>world</sup>'", "hello ^world^", "hello <sup>world</sup>"},
		{"'this ^is^ subscript' should become 'this <sup>is</sup> subscript'", "this ^is^ subscript", "this <sup>is</sup> subscript"},
		{"'^this^ is ^subscript^' should become '<sup>this</sup> is <sup>subscript</sup>'", "^this^ is ^subscript^", "<sup>this</sup> is <sup>subscript</sup>"},
	}

	for _, test := range cases {
		Convey(test.Description, t, func() {
			got := helper.Superscript(test.Input)
			So(got, ShouldEqual, test.Expected)
		})
	}
}
