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

func New(client client.Renderer) *Render {
	return &Render{
		client: client,
		hMutex: &sync.Mutex{},
		jMutex: &sync.Mutex{},
	}
}

// Page resolves the rendering of a specific page with a given model and template name
func (r *Render) Page(w io.Writer, pageModel interface{}, templateName string) {
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
