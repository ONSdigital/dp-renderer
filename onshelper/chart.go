package onshelper

import (
	"bytes"

	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/dp-renderer/client"
)

func (h *Helper) ONSChartResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	model, err := h.getContent(contentPath)
	if err != nil {
		return "", err
	}

	layout := ""
	rc := render.New(client.NewUnrolledAdapterWithLayout(h.asset, h.assetNames, h.development, layout), h.patternLibraryAssetsPath, h.siteDomain)

	buf := new(bytes.Buffer)
	rc.BuildPage(buf, model, "partials/chart")
	return buf.String(), nil
}

func (h *Helper) ONSTableResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	model, err := h.getContent(contentPath)
	if err != nil {
		return "", err
	}

	layout := ""
	rc := render.New(client.NewUnrolledAdapterWithLayout(h.asset, h.assetNames, h.development, layout), h.patternLibraryAssetsPath, h.siteDomain)

	buf := new(bytes.Buffer)
	rc.BuildPage(buf, model, "partials/table")
	return buf.String(), nil
}
