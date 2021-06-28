package render_test

import (
	"net/http/httptest"
	"testing"

	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/dp-renderer/client"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRenderPageMethod(t *testing.T) {
	Convey("Given the renderer is instantiated with a render client", t, func() {
		mockClient := client.NewMockRenderingClient([]string{"test-article-page", "test-homepage"})
		renderer := render.New(mockClient, "path/to/assets", "site-domain")
		w := httptest.NewRecorder()

		Convey("When the renderer's Page method is called with a valid template name and page model", func() {
			mockPage := renderer.NewBasePageModel()
			renderer.BuildPage(w, mockPage, "test-homepage")

			Convey("Then the render client's HTML method should be called, and the page model should have assets path and site domain injected", func() {
				So(mockClient.ValidBuildHTMLMethodCalls, ShouldEqual, 1)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 0)
				So(mockPage.PatternLibraryAssetsPath, ShouldEqual, "path/to/assets")
				So(mockPage.SiteDomain, ShouldEqual, "site-domain")
			})
		})

		Convey("When the renderer's Page method is called with an invalid template name", func() {
			mockPage := renderer.NewBasePageModel()
			renderer.BuildPage(w, mockPage, "")

			Convey("Then the render client should set an error", func() {
				So(mockClient.ValidBuildHTMLMethodCalls, ShouldEqual, 0)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 1)
			})
		})
	})
}
