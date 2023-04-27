package helper_test

import (
	"strings"
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func mockAssetFunction(name string) ([]byte, error) {
	if strings.Contains(name, ".cy.toml") {
		return []byte("[Foo]\ndescription = \"Used for localisation tests\"\nzero = \"No Foos (Welsh)\"\none = \"One Foo (Welsh)\"\ntwo = \"Two Foos (Welsh)\"\nfew = \"Three Foos (Welsh)\"\nmany = \"Six Foos (Welsh)\"\nother = \"Four, five or more than six but not six Foos (Welsh)\""), nil
	}
	return []byte("[Foo]\ndescription = \"Used for localisation tests\"\none = \"One Foo (English)\"\nother = \"Two or more Foos (English)\""), nil
}

// TestLocalise ensures that the correct strings are returned given a range of different arguments
func TestLocalise(t *testing.T) {
	helper.InitialiseLocalisationsHelper(mockAssetFunction)
	Convey("English localisation", t, func() {
		Convey("English singular is returned", func() {
			So(helper.Localise("Foo", "en", 1), ShouldEqual, "One Foo (English)")
		})
		Convey("English plural is returned for more than one", func() {
			So(helper.Localise("Foo", "en", 4), ShouldEqual, "Two or more Foos (English)")
		})
	})

	Convey("Welsh localisation", t, func() {
		Convey("Welsh nil value sentence returned", func() {
			So(helper.Localise("Foo", "cy", 0), ShouldEqual, "No Foos (Welsh)")
		})
		Convey("Welsh singular is returned", func() {
			So(helper.Localise("Foo", "cy", 1), ShouldEqual, "One Foo (Welsh)")
		})
		Convey("Welsh plural for two is returned", func() {
			So(helper.Localise("Foo", "cy", 2), ShouldEqual, "Two Foos (Welsh)")
		})
		Convey("Welsh plural for few(3) is returned", func() {
			So(helper.Localise("Foo", "cy", 3), ShouldEqual, "Three Foos (Welsh)")
		})
		Convey("Welsh plural for other (4) is returned", func() {
			So(helper.Localise("Foo", "cy", 4), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
		})
		Convey("Welsh plural for other (5) is returned", func() {
			So(helper.Localise("Foo", "cy", 5), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
		})
		Convey("Welsh plural for many (6) is returned", func() {
			So(helper.Localise("Foo", "cy", 6), ShouldEqual, "Six Foos (Welsh)")
		})
		Convey("Welsh plural for many (7) is returned", func() {
			So(helper.Localise("Foo", "cy", 7), ShouldEqual, "Four, five or more than six but not six Foos (Welsh)")
		})
	})
}
