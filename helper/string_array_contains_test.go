package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStringArrayContains(t *testing.T) {
	Convey("That true is returned if the array contains the given comparitor", t, func() {
		stringArray := []string{"cat", "dog"}
		got := helper.StringArrayContains("cat", stringArray)
		want := true
		So(got, ShouldEqual, want)
	})

	Convey("That false is returned if the array does not congain the given comparitor", t, func() {
		stringArray := []string{"cat", "dog"}
		got := helper.StringArrayContains("bat", stringArray)
		want := false
		So(got, ShouldEqual, want)
	})
}
