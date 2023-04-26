package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConcatenateStrings(t *testing.T) {
	Convey("That the returned value is https://www.ons.gov.uk/datasets/cpih01 for filterable pages", t, func() {
		got := helper.ConcatenateStrings("www.ons.gov.uk", "/datasets/", "cpih01")
		want := "www.ons.gov.uk/datasets/cpih01"
		So(got, ShouldEqual, want)
	})
}
