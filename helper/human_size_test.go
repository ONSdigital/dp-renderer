package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHumanSize(t *testing.T) {
	Convey("humanSize should return the expected value for 1500 bytes", t, func() {
		res, err := helper.HumanSize("1500")
		So(err, ShouldBeNil)
		So(res, ShouldEqual, "1.5 KB")
	})

	Convey("humanSize should return the expected value for empty input", t, func() {
		res, err := helper.HumanSize("")
		So(err, ShouldBeNil)
		So(res, ShouldBeEmpty)
	})

	Convey("humanSize should return error for non numeric input", t, func() {
		res, err := helper.HumanSize("green eggs and ham")
		So(err, ShouldNotBeNil)
		So(res, ShouldBeEmpty)
	})
}
