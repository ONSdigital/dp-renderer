package sixteenstagresolver

import (
	"errors"
	"strings"

	"github.com/ONSdigital/dp-renderer/model"
)

type ResourceReader struct {
	GetFigure       func(uri string) (model.Figure, error)
	GetResourceBody func(uri string) ([]byte, error)
	GetFileSize     func(uri string) (int, error)
	GetTable        func(html []byte) (string, error)
}

type resourceReader struct {
	reader ResourceReader
	uri    string
}

func (r *resourceReader) getPathUri(path string) string {
	if !strings.HasPrefix(path, r.uri) {
		return r.uri + "/" + path
	}
	return path
}

func (r *resourceReader) GetFigure(path string) (model.Figure, error) {
	if r.reader.GetFigure == nil {
		return model.Figure{}, errors.New("Invalid resource reader. Missing GetFigure func")
	}
	return r.reader.GetFigure(r.getPathUri(path))
}

func (r *resourceReader) GetResourceBody(path string) ([]byte, error) {
	if r.reader.GetResourceBody == nil {
		return nil, errors.New("Invalid resource reader. Missing GetResourceBody func")
	}
	return r.reader.GetResourceBody(r.getPathUri(path))
}

func (r *resourceReader) GetFileSize(path string) (int, error) {
	if r.reader.GetFileSize == nil {
		return 0, errors.New("Invalid resource reader. Missing GetFileSize func")
	}
	return r.reader.GetFileSize(r.getPathUri(path))
}

func (r *resourceReader) GetTable(html []byte) (string, error) {
	return r.reader.GetTable(html)
}
