package render

import (
	"context"
	"io"
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

func New(client client.Renderer, assetsPath, siteDomain string) *Render {
	return &Render{
		client:                   client,
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
		r.error(w, 500, model.ErrorResponse{
			Error: err.Error(),
		})
		log.Event(ctx, "failed to render template", log.Error(err), log.ERROR)
		return
	}
	log.Event(ctx, "rendered template", log.Data{"template": templateName}, log.INFO)
}

// NewBasePageModel wraps around the model package's NewPage function, but injects the assets path and site domain from the render struct.
// This is to negate the need for the caller to have to provide these values for every new page created in a frontend service
func (r *Render) NewBasePageModel() *model.Page {
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
