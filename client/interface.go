package client

import (
	"html/template"
	"io"

	"github.com/ONSdigital/dp-renderer/model"
)

type Renderer interface {
	BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}, funcMap template.FuncMap) error
	SetError(w io.Writer, status int, errorModel model.ErrorResponse) error
}
