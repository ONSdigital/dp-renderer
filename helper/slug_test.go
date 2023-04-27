package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSlug(t *testing.T) {
	cases := []struct {
		Description string
		Input       string
		Expected    string
	}{
		{"'hello world' should become 'hello-world'", "hello world", "hello-world"},
		{"'The Quick Brown Fox Jumps Over The Lazy Dog' should become 'the-quick-brown-fox-jumps-over-the-lazy-dog'",
			"The Quick Brown Fox Jumps Over The Lazy Dog", "the-quick-brown-fox-jumps-over-the-lazy-dog"},
	}

	for _, test := range cases {
		Convey(test.Description, t, func() {
			got := helper.Slug(test.Input)
			So(got, ShouldEqual, test.Expected)
		})
	}
}
