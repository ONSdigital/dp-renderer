package tagresolver

import (
	"bytes"
	"strconv"

	"github.com/ONSdigital/dp-renderer/model"
)

func (h *TagResolverHelper) ONSBoxResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	// align := match[1]   // align attribute
	content := match[2] // tag content

	return h.applyTemplate(model.Figure{Content: content}, "partials/ons-box"), nil
}

func (h *TagResolverHelper) ONSChartResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	return h.getContentAndApplyTemplate(h.resourceReader.GetFigure, contentPath, "partials/ons-chart")
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
			resource, err := h.resourceReader.GetResourceBody(contentPath + "." + sidecarFile.FileType)
			if err != nil {
				return "", err
			}
			figure.Files[i].Content = string(resource)
		}
	}

	return h.applyTemplate(figure, "partials/ons-equation"), nil
}

func (h *TagResolverHelper) ONSImageResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}
	for i, sidecarFile := range figure.Files {
		size, err := h.resourceReader.GetFileSize(contentPath + "." + sidecarFile.FileType)
		if err != nil {
			return "", err
		}
		// TODO convert size to human readable byte count.
		// See https://github.com/ONSdigital/babbage/blob/c77bf4936a4c8872c674e974a3e9c08d1ad89cf4/src/main/java/com/github/onsdigital/babbage/template/handlebars/helpers/resolve/DataHelpers.java#L276-L282
		figure.Files[i].FileSize = strconv.Itoa(size)
	}

	return h.applyTemplate(figure, "partials/ons-image"), nil
}

func (h *TagResolverHelper) ONSQuoteResolver(match []string) (string, error) {
	model := model.Figure{}
	// figureTag := match[0]   // figure tag
	model.Content = match[1] // content attribute
	if len(match) > 2 {
		model.Attribution = match[2] // attr attribute
	}

	return h.applyTemplate(model, "partials/ons-quote"), nil
}

func (h *TagResolverHelper) ONSTableResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	return h.getContentAndApplyTemplate(h.resourceReader.GetFigure, contentPath, "partials/ons-table")
}

func (h *TagResolverHelper) ONSTable2Resolver(match []string) (string, error) {
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

	return h.applyTemplate(figure, "partials/ons-tablev2"), nil
}

func (h *TagResolverHelper) ONSWarningResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	content := match[1] // tag content

	return h.applyTemplate(model.Figure{Content: content}, "partials/ons-warning"), nil
}

func (h *TagResolverHelper) applyTemplate(figure interface{}, template string) string {
	buf := new(bytes.Buffer)
	h.render.BuildPage(buf, figure, template)
	return buf.String()
}

func (h *TagResolverHelper) getContentAndApplyTemplate(getContent func(string) (model.Figure, error), path string, template string) (string, error) {
	model, err := getContent(path)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	h.render.BuildPage(buf, model, template)
	return buf.String(), nil
}
