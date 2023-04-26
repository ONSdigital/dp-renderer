package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSubtract(t *testing.T) {
	Convey("substract should return expected value", t, func() {
		So(helper.Subtract(100, 1), ShouldEqual, 99)
	})
}
