package onshelper

import (
	"bytes"
	"encoding/json"

	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/dp-renderer/client"
)

type Mapper struct {
	Title string
	Axis  int
}

func (h *Helper) ONSChartResolver(getContent GetContent, match []string) (string, error) {

	contentPath := match[1]

	content, err := h.getContent(contentPath) 

	if err != nil {
		return "", err
	}

	model, err := buildGenericModel(content)

	if err != nil {
		return "", err
	}
	
	rc := render.New(client.NewUnrolledAdapterWithPartials(h.asset, h.assetNames, h.development), h.patternLibraryAssetsPath, h.siteDomain)

	buf := new(bytes.Buffer)
	rc.BuildPage(buf, model, "partials/chart")

	return buf.String(), nil
}

func buildGenericModel(b []byte) (*Mapper, error) {
	var mapper Mapper
	if err := json.Unmarshal(b, &mapper); err != nil {
		return nil, err
	}

	return &mapper, nil
}
