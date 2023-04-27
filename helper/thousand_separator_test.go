package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestThousandsSeparator(t *testing.T) {
	Convey("Given an integer to separate", t, func() {
		Convey("When a valid integer is passed", func() {
			Convey("Then a string of comma separated thousands is returned", func() {
				So(helper.ThousandsSeparator(1), ShouldEqual, "1")
				So(helper.ThousandsSeparator(123), ShouldEqual, "123")
				So(helper.ThousandsSeparator(1234), ShouldEqual, "1,234")
				So(helper.ThousandsSeparator(12345), ShouldEqual, "12,345")
				So(helper.ThousandsSeparator(123456), ShouldEqual, "123,456")
				So(helper.ThousandsSeparator(1234567), ShouldEqual, "1,234,567")
				So(helper.ThousandsSeparator(1234567890), ShouldEqual, "1,234,567,890")
				So(helper.ThousandsSeparator(-1234567890), ShouldEqual, "-1,234,567,890")
			})
		})
	})
}
