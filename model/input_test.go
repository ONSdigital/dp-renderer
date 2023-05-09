package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFuncGetInputType(t *testing.T) {
	Convey("Given a text input field", t, func() {
		Convey("When the FuncGetInputType function is called", func() {
			Convey("Then it returns the expected input type", func() {
				tc := []struct {
					given    int
					expected string
				}{
					{
						// text returned to ensure backwards compatibility
						given:    0,
						expected: "text",
					},
					{
						given:    int(Text),
						expected: "text",
					},
					{
						given:    int(Email),
						expected: "email",
					},
					{
						given:    int(Tel),
						expected: "tel",
					},
					{
						given:    int(Url),
						expected: "url",
					},
					{
						given:    25,
						expected: "",
					},
				}
				for _, t := range tc {
					mockType := Input{
						Type: InputType(t.given),
					}
					So(mockType.FuncGetInputType(), ShouldEqual, t.expected)
				}
			})
		})
	})
}
