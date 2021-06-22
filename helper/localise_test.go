package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/assets"
	"github.com/ONSdigital/dp-renderer/helper"
	. "github.com/smartystreets/goconvey/convey"
)

// TestLocalise ensures that the correct strings are returned given a range of different arguments
func TestLocalise(t *testing.T) {
	helper.InitialiseLocalisationsHelper(assets.Asset)
	Convey("English singular is returned", t, func() {
		So(helper.Localise("Foo", "en", 1), ShouldEqual, "One Foo (English)")
	})
	Convey("English plural is returned for more than one", t, func() {
		So(helper.Localise("Foo", "en", 4), ShouldEqual, "Two or more Foos (English)")
	})
	Convey("Welsh nil value sentence returned", t, func() {
		So(helper.Localise("Foo", "cy", 0), ShouldEqual, "No Foos (Welsh)")
	})
	Convey("Welsh singular is returned", t, func() {
		So(helper.Localise("Foo", "cy", 1), ShouldEqual, "One Foo (Welsh)")
	})
	Convey("Welsh plural for two is returned", t, func() {
		So(helper.Localise("Foo", "cy", 2), ShouldEqual, "Two Foos (Welsh)")
	})
	Convey("Welsh plural for few(3) is returned", t, func() {
		So(helper.Localise("Foo", "cy", 3), ShouldEqual, "Three Foos (Welsh)")
	})
	Convey("Welsh plural for other (4) is returned", t, func() {
		So(helper.Localise("Foo", "cy", 4), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
	})
	Convey("Welsh plural for other (5) is returned", t, func() {
		So(helper.Localise("Foo", "cy", 5), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
	})
	Convey("Welsh plural for many (6) is returned", t, func() {
		So(helper.Localise("Foo", "cy", 6), ShouldEqual, "Six Foos (Welsh)")
	})
	Convey("Welsh plural for many (7) is returned", t, func() {
		So(helper.Localise("Foo", "cy", 7), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
	})
}
