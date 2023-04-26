package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTruncateToMaximumCharacters(t *testing.T) {
	Convey("That text is not truncated when it contains fewer characters than maxLength", t, func() {
		got := helper.TruncateToMaximumCharacters("Hello world", 50)
		want := "Hello world"
		So(got, ShouldEqual, want)
	})
	Convey("That text is truncated when exceeding maxLength", t, func() {
		got := helper.TruncateToMaximumCharacters("Hello world", 2)
		want := "He..."
		So(got, ShouldEqual, want)
	})
	Convey("That leading/trailling whitespace is removed when text is truncated", t, func() {
		got := helper.TruncateToMaximumCharacters("The space after 'space' should not be included", 10)
		want := "The space..."
		So(got, ShouldEqual, want)
	})
}
