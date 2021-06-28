package render_test

import (
	"net/http/httptest"
	"testing"

	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/dp-renderer/client"
	"github.com/ONSdigital/dp-renderer/model"
	. "github.com/smartystreets/goconvey/convey"
)

func mockAssetFunction(name string) ([]byte, error) {
	return []byte(""), nil
}

func TestRenderPageMethod(t *testing.T) {
	Convey("Given the renderer is instantiated with a render client", t, func() {
		mockClient := client.NewMockRenderingClient([]string{"test-article-page", "test-homepage"})
		renderer := render.New(mockClient)
		w := httptest.NewRecorder()

		Convey("When the renderer's Page method is called with a valid template name and page model", func() {
			mockPage := model.Page{}
			renderer.Page(w, mockPage, "test-homepage")

			Convey("Then the render client's HTML method should be called", func() {
				So(mockClient.ValidBuildHTMLMethodCalls, ShouldEqual, 1)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 0)
			})
		})

		Convey("When the renderer's Page method is called with an invalid template name", func() {
			mockPage := model.Page{
				PatternLibraryAssetsPath: "path/to/assets",
			}
			renderer.Page(w, mockPage, "")

			Convey("Then the render client should set an error", func() {
				So(mockClient.ValidBuildHTMLMethodCalls, ShouldEqual, 0)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 1)
			})
		})
	})
}
