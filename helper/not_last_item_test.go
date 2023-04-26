package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNotLastItem(t *testing.T) {
	Convey("That true is returned when index is 1 and length is 3", t, func() {
		got := helper.NotLastItem(3, 1)
		So(got, ShouldEqual, true)
	})
	Convey("That false is returned when index is 0 and length is 1", t, func() {
		got := helper.NotLastItem(1, 0)
		So(got, ShouldEqual, false)
	})
}
