package onshelper

import (
	"bytes"
)

func (h *Helper) ONSChartResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	return h.getContentAndApplyTemplate(contentPath, "partials/chart")
}

func (h *Helper) ONSTableResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	return h.getContentAndApplyTemplate(contentPath, "partials/table")
}

func (h *Helper) getContentAndApplyTemplate(path string, template string) (string, error) {
	model, err := h.getContent(path)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	h.render.BuildPage(buf, model, template)
	return buf.String(), nil
}
