package tagresolver

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/ONSdigital/dp-renderer/model"
)

func (h *TagResolverHelper) ONSBoxResolver(match []string) (string, error) {
	model := model.Figure{}
	// figureTag := match[0]   // figure tag
	model.Align = match[1]   // align attribute
	model.Content = match[2] // tag content

	return h.applyTemplate(model, "partials/ons-tags/ons-box"), nil
}

func (h *TagResolverHelper) ONSChartResolver(match []string) (string, error) {
	if h.resourceReader.GetFigure == nil {
		return "", errors.New("Invalid resource reader for chart resolver")
	}
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	uri := h.resourceReader.getPathUri(contentPath)
	figure, err := h.resourceReader.GetFigure(uri)
	if err != nil {
		return "", err
	}
	return h.applyTemplate(figure, "partials/ons-tags/ons-chart"), nil
}

func (h *TagResolverHelper) ONSEquationResolver(match []string) (string, error) {
	if h.resourceReader.GetResourceBody == nil {
		return "", errors.New("Invalid resource reader for equation resolver")
	}
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	uri := h.resourceReader.getPathUri(contentPath)
	figure, err := h.resourceReader.GetFigure(uri)
	if err != nil {
		return "", err
	}

	for i, sidecarFile := range figure.Files {
		if sidecarFile.Type == "generated-svg" {
			resource, err := h.resourceReader.GetResourceBody(uri + "." + sidecarFile.FileType)
			if err != nil {
				return "", err
			}
			figure.Files[i].Content = string(resource)
		}
	}

	return h.applyTemplate(figure, "partials/ons-tags/ons-equation"), nil
}

func (h *TagResolverHelper) ONSImageResolver(match []string) (string, error) {
	if h.resourceReader.GetFigure == nil || h.resourceReader.GetFileSize == nil {
		return "", errors.New("Invalid resource reader for image resolver")
	}
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	uri := h.resourceReader.getPathUri(contentPath)
	figure, err := h.resourceReader.GetFigure(uri)
	if err != nil {
		return "", err
	}
	for i, sidecarFile := range figure.Files {
		size, err := h.resourceReader.GetFileSize(uri + "." + sidecarFile.FileType)
		if err != nil {
			return "", err
		}
		figure.Files[i].FileSize = humanReadableByteCount(size)
	}

	return h.applyTemplate(figure, "partials/ons-tags/ons-image"), nil
}

func humanReadableByteCount(b int) string {
	var unit float64 = 1000
	bytes := float64(b)
	if bytes < unit {
		return strconv.Itoa(b) + " B"
	}
	exp := (int)(math.Log(bytes) / math.Log(unit))
	pre := string("kMGTPE"[exp-1])
	return fmt.Sprintf("%.1f %sB", bytes/math.Pow(unit, float64(exp)), pre)
}

func (h *TagResolverHelper) ONSQuoteResolver(match []string) (string, error) {
	model := model.Figure{}
	// figureTag := match[0]   // figure tag
	model.Content = match[1] // content attribute
	if len(match) > 2 {
		model.Attribution = match[2] // attr attribute
	}

	return h.applyTemplate(model, "partials/ons-tags/ons-quote"), nil
}

func (h *TagResolverHelper) ONSTableResolver(match []string) (string, error) {
	if h.resourceReader.GetFigure == nil {
		return "", errors.New("Invalid resource reader for table resolver")
	}
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	uri := h.resourceReader.getPathUri(contentPath)
	figure, err := h.resourceReader.GetFigure(uri)
	if err != nil {
		return "", err
	}

	for i, sidecarFile := range figure.Files {
		if sidecarFile.Type == "html" {
			resource, err := h.resourceReader.GetResourceBody(uri + ".html")
			if err != nil {
				return "", err
			}
			figure.Files[i].Content = string(resource)
		}
	}

	fmt.Printf("ONSTableResolver() figure %#v", figure)
	return h.applyTemplate(figure, "partials/ons-tags/ons-table"), nil
}

func (h *TagResolverHelper) ONSTableV2Resolver(match []string) (string, error) {
	if h.resourceReader.GetResourceBody == nil ||
		h.resourceReader.GetTable == nil ||
		h.resourceReader.GetFigure == nil {
		return "", errors.New("Invalid resource reader for table v2 resolver")
	}
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	uri := h.resourceReader.getPathUri(contentPath)
	html, err := h.resourceReader.GetResourceBody(uri + ".json")
	if err != nil {
		return "", err
	}

	table, err := h.resourceReader.GetTable(html)
	if err != nil {
		return "", err
	}

	figure, err := h.resourceReader.GetFigure(uri)
	if err != nil {
		return "", err
	}
	figure.Content = table

	return h.applyTemplate(figure, "partials/ons-tags/ons-table-v2"), nil
}

func (h *TagResolverHelper) ONSWarningResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	content := match[1] // tag content

	return h.applyTemplate(model.Figure{Content: content}, "partials/ons-tags/ons-warning"), nil
}

func (h *TagResolverHelper) applyTemplate(figure interface{}, template string) string {
	buf := new(bytes.Buffer)
	h.render.BuildPage(buf, figure, template)
	return buf.String()
}
