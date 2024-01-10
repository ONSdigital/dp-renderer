package render_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	render "github.com/ONSdigital/dp-renderer/v2"
	"github.com/ONSdigital/dp-renderer/v2/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRenderPageMethod(t *testing.T) {
	Convey("Given the renderer is instantiated with a render client", t, func() {
		mockClient := newMockRenderingClient([]string{"test-article-page", "test-homepage", "error/401", "error/404", "error/500"})
		renderer := render.New(mockClient, "path/to/assets", "site-domain")
		w := httptest.NewRecorder()

		Convey("When the renderer's Page method is called with a valid template name and page model", func() {
			mockPage := renderer.NewBasePageModel()
			renderer.BuildPage(w, mockPage, "test-homepage")

			Convey("Then the render client's HTML method should be called, and the page model should have assets path and site domain injected", func() {
				So(len(mockClient.ValidBuildHTMLMethodCalls), ShouldEqual, 1)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 0)
				So(mockPage.PatternLibraryAssetsPath, ShouldEqual, "path/to/assets")
				So(mockPage.SiteDomain, ShouldEqual, "site-domain")
			})
		})

		Convey("When the renderer's Page method is called with an invalid template name", func() {
			mockPage := renderer.NewBasePageModel()
			renderer.BuildPage(w, mockPage, "")

			Convey("Then the render client should set an error", func() {
				So(mockClient.ValidBuildHTMLMethodCalls, ShouldBeEmpty)
				So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 1)
			})
		})

		Convey("When the renderer's build error page method is called", func() {
			mockPage := renderer.NewBasePageModel()

			Convey("Then the render client should set pageModel and call build HTML func", func() {
				Convey("for 401", func() {
					renderer.BuildErrorPage(w, mockPage, http.StatusUnauthorized)
					expectedPageModel := mockClient.ValidBuildHTMLMethodCalls[0].PageModel.(model.Page)
					So(mockClient.ValidBuildHTMLMethodCalls, ShouldHaveLength, 1)
					So(expectedPageModel.Error.Title, ShouldEqual, "Access denied")
					So(expectedPageModel.Error.ErrorCode, ShouldEqual, http.StatusUnauthorized)
					So(expectedPageModel.Enable500ErrorPageStyling, ShouldBeFalse)
					So(mockClient.ValidBuildHTMLMethodCalls[0].TemplateName, ShouldEqual, "error/401")
					So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 0)
				})
				Convey("for 404", func() {
					renderer.BuildErrorPage(w, mockPage, http.StatusNotFound)
					expectedPageModel := mockClient.ValidBuildHTMLMethodCalls[0].PageModel.(model.Page)
					So(mockClient.ValidBuildHTMLMethodCalls, ShouldHaveLength, 1)
					So(expectedPageModel.Error.Title, ShouldEqual, "Page not found")
					So(expectedPageModel.Error.ErrorCode, ShouldEqual, http.StatusNotFound)
					So(expectedPageModel.Enable500ErrorPageStyling, ShouldBeFalse)
					So(mockClient.ValidBuildHTMLMethodCalls[0].TemplateName, ShouldEqual, "error/404")
					So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 0)
				})
				Convey("for 500", func() {
					renderer.BuildErrorPage(w, mockPage, http.StatusInternalServerError)
					expectedPageModel := mockClient.ValidBuildHTMLMethodCalls[0].PageModel.(model.Page)
					So(mockClient.ValidBuildHTMLMethodCalls, ShouldHaveLength, 1)
					So(expectedPageModel.Error.Title, ShouldEqual, "Sorry, there's a problem with the service")
					So(expectedPageModel.Error.ErrorCode, ShouldEqual, http.StatusInternalServerError)
					So(expectedPageModel.Enable500ErrorPageStyling, ShouldBeTrue)
					So(mockClient.ValidBuildHTMLMethodCalls[0].TemplateName, ShouldEqual, "error/500")
					So(mockClient.ValidSetErrorMethodCalls, ShouldEqual, 0)
				})
			})
		})
	})
}

type mockRenderClient struct {
	TemplateNames             []string
	ValidBuildHTMLMethodCalls []ValidBuildHTMLMethod
	ValidSetErrorMethodCalls  int
}
type ValidBuildHTMLMethod struct {
	Status       int
	TemplateName string
	PageModel    interface{}
}

func newMockRenderingClient(templateNames []string) *mockRenderClient {
	return &mockRenderClient{
		TemplateNames: templateNames,
	}
}

func (m *mockRenderClient) BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error {
	for _, value := range m.TemplateNames {
		if value == templateName {
			call := ValidBuildHTMLMethod{Status: status, TemplateName: templateName, PageModel: pageModel}
			m.ValidBuildHTMLMethodCalls = append(m.ValidBuildHTMLMethodCalls, call)
			return nil
		}
	}
	return errors.New("Failed to build page")
}

func (m *mockRenderClient) SetError(w io.Writer, status int, errorModel model.ErrorResponse) error {
	m.ValidSetErrorMethodCalls++
	return errors.New("An error occurred when building the page")
}
