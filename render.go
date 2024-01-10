package render

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ONSdigital/dp-renderer/v2/client"
	"github.com/ONSdigital/dp-renderer/v2/model"
	"github.com/ONSdigital/log.go/v2/log"
)

type Render struct {
	client                               client.Renderer
	hMutex                               *sync.Mutex
	jMutex                               *sync.Mutex
	PatternLibraryAssetsPath, SiteDomain string
}

// New returns a render struct and accepts any rendering client that satisfies the Renderer interface
func New(client client.Renderer, assetsPath, siteDomain string) *Render {
	return &Render{
		client:                   client,
		hMutex:                   &sync.Mutex{},
		jMutex:                   &sync.Mutex{},
		PatternLibraryAssetsPath: assetsPath,
		SiteDomain:               siteDomain,
	}
}

/*
NewWithDefaultClient returns a render struct with a default rendering client provided (default: unrolled/render)
When the siteDomain argument contains "localhost", then the rendering client will be instantiated in "development" mode.
This means that templates are recompiled on request.
Any updates made to your templates can then be viewed upon browser refresh, rather than having to restart the app.
*/
func NewWithDefaultClient(assetFn func(name string) ([]byte, error), assetNameFn func() []string, assetsPath, siteDomain string) *Render {
	isDevelopment := false
	if strings.Contains(siteDomain, "localhost") {
		isDevelopment = true
	}
	return &Render{
		client:                   client.NewUnrolledAdapter(assetFn, assetNameFn, isDevelopment),
		hMutex:                   &sync.Mutex{},
		jMutex:                   &sync.Mutex{},
		PatternLibraryAssetsPath: assetsPath,
		SiteDomain:               siteDomain,
	}
}

// BuildPage resolves the rendering of a specific page with a given model and template name
func (r *Render) BuildPage(w io.Writer, pageModel interface{}, templateName string) {
	ctx := context.Background()
	if err := r.render(w, 200, templateName, pageModel); err != nil {
		log.Error(ctx, "failed to render template", err, log.Data{"template": templateName})
		if modelErr := r.error(w, 500, model.ErrorResponse{
			Error: err.Error(),
		}); modelErr != nil {
			log.Error(ctx, "failed to set error response", modelErr)
		}
		return
	}
	log.Info(ctx, "rendered template", log.Data{"template": templateName})
}

// BuildErrorPage resolves the rendering of a specific page with a given model and template name
func (r *Render) BuildErrorPage(w io.Writer, pageModel model.Page, statusCode int) {
	// set template name based on http status code
	templateName := "error/" + strconv.Itoa(statusCode)

	switch statusCode {
	case http.StatusUnauthorized:
		pageModel.Error.Title = "Access denied"
	case http.StatusNotFound:
		pageModel.Error.Title = "Page not found"
	case http.StatusInternalServerError:
		pageModel.Error.Title = "Sorry, there's a problem with the service"
		pageModel.Enable500ErrorPageStyling = true
	}
	pageModel.Error.ErrorCode = statusCode

	ctx := context.Background()
	if err := r.render(w, statusCode, templateName, pageModel); err != nil {
		log.Error(ctx, "failed to render error template", err, log.Data{"template": templateName})
		if modelErr := r.error(w, 500, model.ErrorResponse{
			Error: err.Error(),
		}); modelErr != nil {
			log.Error(ctx, "failed to set error response", modelErr)
		}
		return
	}
	log.Info(ctx, "rendered error template", log.Data{"template": templateName})
}

// NewBasePageModel wraps around the model package's NewPage function, but injects the assets path and site domain from the render struct.
// This is to negate the need for the caller to have to provide these values for every new page created in a frontend service
func (r *Render) NewBasePageModel() model.Page {
	return model.NewPage(r.PatternLibraryAssetsPath, r.SiteDomain)
}

func (r *Render) render(w io.Writer, status int, templateName string, pageModel interface{}) error {
	r.hMutex.Lock()
	defer r.hMutex.Unlock()
	return r.client.BuildHTML(w, status, templateName, pageModel)
}

func (r *Render) error(w io.Writer, status int, errorModel model.ErrorResponse) error {
	r.jMutex.Lock()
	defer r.jMutex.Unlock()
	return r.client.SetError(w, status, errorModel)
}
