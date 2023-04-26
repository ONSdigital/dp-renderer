package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLower(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Expected    string
	}{
		{"'HELLO WORLD' should become 'hello world'", "HELLO WORLD", "hello world"},
		{"'Hello World' should become 'hello world'", "Hello World", "hello world"},
		{"'hEllO WoRlD' should become 'hello world'", "hEllO WoRlD", "hello world"},
		{"'hello world' should become 'hello world'", "hello world", "hello world"},
	}

	for _, test := range cases {
		Convey(test.Description, t, func() {
			got := helper.Lower(test.Input)
			So(got, ShouldEqual, test.Expected)
		})
	}
}
