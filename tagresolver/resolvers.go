package tagresolver

import (
	"bytes"
)

func (h *TagResolverHelper) ONSChartResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	return h.getContentAndApplyTemplate(contentPath, "partials/ons-chart")
}

func (h *TagResolverHelper) ONSTableResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	return h.getContentAndApplyTemplate(contentPath, "partials/ons-table")
}

func (h *TagResolverHelper) getContentAndApplyTemplate(path string, template string) (string, error) {
	model, err := h.getContent(path)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	h.render.BuildPage(buf, model, template)
	return buf.String(), nil
}
