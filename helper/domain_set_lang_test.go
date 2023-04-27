package helper_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/helper"
	. "github.com/smartystreets/goconvey/convey"
)

// TestDomainSetLang ensures the returned URL is set accurately
func TestDomainSetLang(t *testing.T) {
	Convey("English domain requested", t, func() {
		So(helper.DomainSetLang("www.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("https://www.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("www.cy.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("cy.ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("https://www.cy.ons.gov.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://cy.ons.gov.uk", "", "en"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("www.foo-bar.baz.co.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(helper.DomainSetLang("cy.foo-bar.baz.co.uk", "", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(helper.DomainSetLang("https://www.cy.foo-bar.baz.co.uk", "/foo/bar/baz", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://cy.foo-bar.baz.co.uk", "", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(helper.DomainSetLang("https://cy.foo-bar.baz.co.uk", "https://foo:12345/bar/baz/qux", "en"), ShouldEqual, "https://www.foo-bar.baz.co.uk/bar/baz/qux")
	})

	Convey("Welsh domain requested", t, func() {
		So(helper.DomainSetLang("www.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(helper.DomainSetLang("ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://www.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(helper.DomainSetLang("https://ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("www.cy.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(helper.DomainSetLang("cy.ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://www.cy.ons.gov.uk", "", "cy"), ShouldEqual, "https://cy.ons.gov.uk")
		So(helper.DomainSetLang("https://cy.ons.gov.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("www.foo-bar.baz.co.uk", "", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk")
		So(helper.DomainSetLang("cy.foo-bar.baz.co.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://www.cy.foo-bar.baz.co.uk", "", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk")
		So(helper.DomainSetLang("https://cy.foo-bar.baz.co.uk", "/foo/bar/baz", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://cy.foo-bar.baz.co.uk", "https://foo:12345/bar/baz/qux", "cy"), ShouldEqual, "https://cy.foo-bar.baz.co.uk/bar/baz/qux")
	})

	Convey("Unsupported domain requested", t, func() {
		So(helper.DomainSetLang("www.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://www.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("https://ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("www.cy.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("cy.ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://www.cy.ons.gov.uk", "", "foo"), ShouldEqual, "https://www.ons.gov.uk")
		So(helper.DomainSetLang("https://cy.ons.gov.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.ons.gov.uk/foo/bar/baz")
		So(helper.DomainSetLang("www.foo-bar.baz.co.uk", "", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(helper.DomainSetLang("cy.foo-bar.baz.co.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://www.cy.foo-bar.baz.co.uk", "", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk")
		So(helper.DomainSetLang("https://cy.foo-bar.baz.co.uk", "/foo/bar/baz", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk/foo/bar/baz")
		So(helper.DomainSetLang("https://cy.foo-bar.baz.co.uk", "https://foo:12345/bar/baz/qux", "foo"), ShouldEqual, "https://www.foo-bar.baz.co.uk/bar/baz/qux")
	})
}
