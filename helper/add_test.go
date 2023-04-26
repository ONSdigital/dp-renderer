package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("add should return expected value", t, func() {
		So(helper.Add(99, 1), ShouldEqual, 100)
	})
}
