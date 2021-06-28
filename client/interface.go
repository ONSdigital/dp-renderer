package client

import (
	"io"

	"github.com/ONSdigital/dp-renderer/model"
)

type Renderer interface {
	BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error
	SetError(w io.Writer, status int, errorModel model.ErrorResponse) error
}
