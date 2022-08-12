package client

import (
	"html/template"
	"io"

	"github.com/ONSdigital/dp-renderer/helper"
	"github.com/ONSdigital/dp-renderer/model"
	"github.com/unrolled/render"
	unrolled "github.com/unrolled/render"
)

type unrolledAdapter struct {
	unrolled *unrolled.Render
}

// NewUnrolledAdapterWithPartials returns the unrolled render library via an adapter struct that satisfies the Renderer interface
func NewUnrolledAdapterWithPartials(assetFn func(name string) ([]byte, error), assetNameFn func() []string, isDevelopment bool) Renderer {
	helper.InitialiseLocalisationsHelper(assetFn)

	return &unrolledAdapter{
		unrolled: unrolled.New(render.Options{
			Asset:         assetFn,
			AssetNames:    assetNameFn,
			Layout:        "",
			IsDevelopment: isDevelopment,
			Funcs:         []template.FuncMap{helper.RegisteredFuncs},
		}),
	}
}

// NewUnrolledAdapter returns the unrolled render library via an adapter struct that satisfies the Renderer interface
func NewUnrolledAdapter(assetFn func(name string) ([]byte, error), assetNameFn func() []string, isDevelopment bool) Renderer {
	helper.InitialiseLocalisationsHelper(assetFn)

	return &unrolledAdapter{
		unrolled: unrolled.New(render.Options{
			Asset:         assetFn,
			AssetNames:    assetNameFn,
			Layout:        "main",
			IsDevelopment: isDevelopment,
			Funcs:         []template.FuncMap{helper.RegisteredFuncs},
		}),
	}
}

// BuildHTML produces the HTML content based on the template and page model provided
func (adapter *unrolledAdapter) BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}, funcMap template.FuncMap) error {
	if funcMap != nil {
		options := unrolled.HTMLOptions{
			Funcs: funcMap,
		}
		return adapter.unrolled.HTML(w, status, templateName, pageModel, options)
	}

	return adapter.unrolled.HTML(w, status, templateName, pageModel)
}

// SetError provides an error response that is mapped to a model
func (adapter *unrolledAdapter) SetError(w io.Writer, status int, errorModel model.ErrorResponse) error {
	return adapter.unrolled.JSON(w, status, errorModel)
}
