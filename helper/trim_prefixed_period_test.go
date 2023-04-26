package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTrimPrefixedPeriod(t *testing.T) {
	Convey("That the returned string is myString", t, func() {
		got := helper.TrimPrefixedPeriod(".myString")
		want := "myString"
		So(got, ShouldEqual, want)
	})
	Convey("That all periods are trimmed", t, func() {
		got := helper.TrimPrefixedPeriod(".....myString")
		want := "myString"
		So(got, ShouldEqual, want)
	})
	Convey("That only leading periods are trimmed", t, func() {
		got := helper.TrimPrefixedPeriod(".string with periods on the end....")
		want := "string with periods on the end...."
		So(got, ShouldEqual, want)
	})
	Convey("That a string with no period is returned", t, func() {
		got := helper.TrimPrefixedPeriod("a string with no period")
		want := "a string with no period"
		So(got, ShouldEqual, want)
	})
}
