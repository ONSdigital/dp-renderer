package render_test

import (
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	render "github.com/ONSdigital/dp-renderer/v2"
	"github.com/ONSdigital/dp-renderer/v2/model"
	"github.com/davecgh/go-spew/spew"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRenderPageMethod(t *testing.T) {
	Convey("Given the renderer is instantiated with a render client", t, func() {
		mockClient := newMockRenderingClient([]string{"test-article-page", "test-homepage"})
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

		Convey("When the renderer's build error page method is called", func() {
			mockPage := renderer.NewBasePageModel()
			renderer.BuildErrorPage(w, mockPage, 401)

			spew.Dump(string(w.Body))

			Convey("Then the render client should set an error", func() {
				So(mockClient.ValidBuildHTMLMethodCalls, ShouldEqual, 2)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 1)
			})
		})
	})
}

type mockRenderClient struct {
	TemplateNames             []string
	ValidBuildHTMLMethodCalls int
	ValidSetErrorMethodCalls  int
}

func newMockRenderingClient(templateNames []string) *mockRenderClient {
	return &mockRenderClient{
		TemplateNames: templateNames,
	}
}

func (m *mockRenderClient) BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error {
	for _, value := range m.TemplateNames {
		if value == templateName {
			m.ValidBuildHTMLMethodCalls++
			return nil
		}
	}
	return errors.New("Failed to build page")
}

func (m *mockRenderClient) SetError(w io.Writer, status int, errorModel model.ErrorResponse) error {
	m.ValidSetErrorMethodCalls++
	return errors.New("An error occurred when building the page")
}
