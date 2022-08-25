package tagresolver

import (
	"bytes"
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
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}
	return h.applyTemplate(figure, "partials/ons-tags/ons-chart"), nil
}

func (h *TagResolverHelper) ONSEquationResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}

	for i, sidecarFile := range figure.Files {
		if sidecarFile.Type == "generated-svg" {
			resource, err := h.resourceReader.GetResourceBody(sidecarFile.Filename)
			if err != nil {
				return "", err
			}
			figure.Files[i].Content = string(resource)
		}
	}

	return h.applyTemplate(figure, "partials/ons-tags/ons-equation"), nil
}

func (h *TagResolverHelper) ONSImageResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}
	for i, sidecarFile := range figure.Files {
		size, err := h.resourceReader.GetFileSize(sidecarFile.Filename)
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
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}

	for i, sidecarFile := range figure.Files {
		if sidecarFile.Type == "html" {
			resource, err := h.resourceReader.GetResourceBody(sidecarFile.Filename)
			if err != nil {
				return "", err
			}
			figure.Files[i].Content = string(resource)
		}

		size, err := h.resourceReader.GetFileSize(sidecarFile.Filename)
		if err != nil {
			return "", err
		}
		figure.Files[i].FileSize = humanReadableByteCount(size)
	}

	fmt.Printf("ONSTableResolver() figure %#v", figure)
	return h.applyTemplate(figure, "partials/ons-tags/ons-table"), nil
}

func (h *TagResolverHelper) ONSTableV2Resolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	html, err := h.resourceReader.GetResourceBody(contentPath + ".json")
	if err != nil {
		return "", err
	}

	table, err := h.resourceReader.GetTable(html)
	if err != nil {
		return "", err
	}

	figure, err := h.resourceReader.GetFigure(contentPath)
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
