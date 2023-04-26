package model_test

import (
	"testing"

	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewPage(t *testing.T) {
	Convey("An instantation of a new page includes the pattern asset library path and site domain provided", t, func() {
		mockPage := model.NewPage("path/to/assets", "site-domain")
		So(mockPage.PatternLibraryAssetsPath, ShouldEqual, "path/to/assets")
		So(mockPage.SiteDomain, ShouldEqual, "site-domain")
	})
}
