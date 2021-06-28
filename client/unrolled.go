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

func NewUnrolledWrapper(assetFn func(name string) ([]byte, error), assetNameFn func() []string) Renderer {
	helper.InitialiseLocalisationsHelper(assetFn)
	return &unrolledAdapter{
		unrolled: unrolled.New(render.Options{
			Asset:      assetFn,
			AssetNames: assetNameFn,
			Layout:     "main",
			Funcs:      []template.FuncMap{helper.RegisteredFuncs},
		}),
	}
}

func (adapter *unrolledAdapter) BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error {
	return adapter.unrolled.HTML(w, status, templateName, pageModel)
}

func (adapter *unrolledAdapter) SetError(w io.Writer, status int, errorModel model.ErrorResponse) error {
	return adapter.unrolled.JSON(w, status, errorModel)
}
