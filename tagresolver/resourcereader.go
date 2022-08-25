package tagresolver

import (
	"strings"

	"github.com/ONSdigital/dp-renderer/model"
)

type ResourceReader struct {
	URI             string
	GetFigure       func(path string) (model.Figure, error)
	GetResourceBody func(path string) ([]byte, error)
	GetTable        func(html []byte) (string, error)
	GetFileSize     func(path string) (int, error)
}

func (r *ResourceReader) getPathUri(path string) string {
	if !strings.HasPrefix(path, r.URI) {
		return r.URI + "/" + path
	}
	return path
}
