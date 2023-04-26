package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiply(t *testing.T) {
	Convey("multiply should return expected value", t, func() {
		So(helper.Multiply(100, 1), ShouldEqual, 100)
	})
}
