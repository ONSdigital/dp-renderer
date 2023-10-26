package renderror

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	render "github.com/ONSdigital/dp-renderer/v2"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRenderr(t *testing.T) {
	Convey("Test renderr middleware runs correctly", t, func() {
		Convey("when getting a 200 status response", func() {
			r, mockedRC := setupTest()
			req, _ := http.NewRequest("GET", "/success", http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, 200)
			So(mockedRC.calls.BuildHTML, ShouldBeEmpty)
		})

		Convey("when getting a 401 status response ", func() {
			r, mockedRC := setupTest()
			req, _ := http.NewRequest("GET", "/unauthorised", http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, 401)
			So(mockedRC.calls.BuildHTML, ShouldHaveLength, 1)
			So(mockedRC.calls.BuildHTML[0].TemplateName, ShouldEqual, "error/401")
		})

		Convey("when getting a 404 status response ", func() {
			r, mockedRC := setupTest()
			req, _ := http.NewRequest("GET", "/not-found", http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, 404)
			So(mockedRC.calls.BuildHTML, ShouldHaveLength, 1)
			So(mockedRC.calls.BuildHTML[0].TemplateName, ShouldEqual, "error/404")
		})

		Convey("when getting a 500 status response ", func() {
			r, mockedRC := setupTest()
			req, _ := http.NewRequest("GET", "/internal-server-error", http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, 500)
			So(mockedRC.calls.BuildHTML, ShouldHaveLength, 1)
			So(mockedRC.calls.BuildHTML[0].TemplateName, ShouldEqual, "error/500")
		})

		Convey("when getting an error and JSON response ", func() {
			r, mockedRC := setupTest()
			req, _ := http.NewRequest("GET", "/json", http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			So(mockedRC.calls.BuildHTML, ShouldBeEmpty)
		})

	})
}

func setupTest() (http.Handler, *RenderClientMock) {
	router := mux.NewRouter()
	mockedRendererClient := &RenderClientMock{
		BuildHTMLFunc: func(w io.Writer, status int, templateName string, pageModel interface{}) error { return nil },
	}
	rendC := render.New(mockedRendererClient, "", "")
	middleware := []alice.Constructor{
		Handler(rendC),
	}
	testAlice := alice.New(middleware...).Then(router)

	router.HandleFunc("/success", func(w http.ResponseWriter, req *http.Request) {})
	router.HandleFunc("/unauthorised", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	})
	router.HandleFunc("/not-found", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})
	router.HandleFunc("/internal-server-error", func(w http.ResponseWriter, req *http.Request) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	})
	router.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
	})

	return testAlice, mockedRendererClient
}
