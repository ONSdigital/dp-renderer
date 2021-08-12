package render

import (
	"context"
	"io"
	"strings"
	"sync"

	"github.com/ONSdigital/dp-renderer/client"
	"github.com/ONSdigital/dp-renderer/model"
	"github.com/ONSdigital/log.go/log"
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

// NewWithDefaultClient returns a render struct with a default rendering client provided (default: unrolled/render)
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
		err = r.error(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.Event(ctx, "failed to render template", log.ERROR, log.Error(err), log.Data{"template": templateName})
		return
	}
	log.Event(ctx, "rendered template", log.INFO, log.Data{"template": templateName})
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
