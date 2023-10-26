package renderror

import (
	"net/http"
	"strings"

	"github.com/ONSdigital/dp-cookies/cookies"
	render "github.com/ONSdigital/dp-renderer/v2"
	"github.com/ONSdigital/log.go/v2/log"
)

type httpResponseInterceptor struct {
	http.ResponseWriter
	req            *http.Request
	intercepted    bool
	headersWritten bool
	headerCache    http.Header
	renderClient   *render.Render
}

func (rI *httpResponseInterceptor) WriteHeader(status int) {
	if status >= 400 && !strings.HasPrefix(rI.Header().Get("Content-Type"), "application/json") {
		log.Info(rI.req.Context(), "Intercepted error response", log.Data{"status": status})
		rI.intercepted = true
		if status == 401 || status == 404 || status == 500 {
			rI.renderErrorPage(status)
		}
	}
	rI.writeHeaders()
	rI.ResponseWriter.WriteHeader(status)
}

func (rI *httpResponseInterceptor) renderErrorPage(code int) {
	m := rI.renderClient.NewBasePageModel()

	// add cookie preferences to error page model
	preferencesCookie := cookies.GetCookiePreferences(rI.req)
	m.CookiesPreferencesSet = preferencesCookie.IsPreferenceSet
	m.CookiesPolicy.Essential = preferencesCookie.Policy.Essential
	m.CookiesPolicy.Usage = preferencesCookie.Policy.Usage

	rI.renderClient.BuildErrorPage(rI.ResponseWriter, m, code)
}

func (rI *httpResponseInterceptor) Write(b []byte) (int, error) {
	if rI.intercepted {
		return len(b), nil
	}
	rI.writeHeaders()
	return rI.ResponseWriter.Write(b)
}

func (rI *httpResponseInterceptor) writeHeaders() {

	if rI.headersWritten {
		return
	}

	for k, v := range rI.headerCache {
		for _, v2 := range v {
			rI.ResponseWriter.Header().Add(k, v2)
		}
	}

	rI.headersWritten = true
}

func (rI *httpResponseInterceptor) Header() http.Header {
	return rI.headerCache
}

// Handler is middleware that renders error pages based on response status codes
func Handler(rendC *render.Render) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			h.ServeHTTP(&httpResponseInterceptor{w, req, false, false, make(http.Header), rendC}, req)
		})
	}
}
