package onshelper

import (
	"html/template"
	"regexp"
	"strings"
	"sync"

	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/dp-renderer/client"
	"github.com/ONSdigital/dp-renderer/helper"
)

type RenderContent func(match []string) (string, error)
type GetContent func(path string) (interface{}, error)

type contentResolver struct {
	Regexp        regexp.Regexp
	RenderContent RenderContent
}

type Helper struct {
	contentResolvers []contentResolver
	getContent       GetContent
	render           *render.Render
}

func NewHelper(asset func(name string) ([]byte, error), assetNames func() []string, patternLibraryAssetsPath string, siteDomain string, getContent GetContent) *Helper {
	isDevelopment := false
	if strings.Contains(siteDomain, "localhost") {
		isDevelopment = true
	}
	helper := &Helper{
		getContent: getContent,
		render:     render.New(client.NewUnrolledAdapterWithLayout(asset, assetNames, isDevelopment, ""), patternLibraryAssetsPath, siteDomain),
	}

	chartResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-chart path=\"(.*)\" />"),
		RenderContent: helper.ONSChartResolver,
	}

	tableResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-table\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
		RenderContent: helper.ONSTableResolver,
	}

	helper.contentResolvers = []contentResolver{chartResolver, tableResolver}

	return helper
}

func (h *Helper) GetFuncMap() template.FuncMap {
	res := make(template.FuncMap)
	//Copy registered funcs
	for k, v := range helper.RegisteredFuncs {
		res[k] = v
	}
	// Override markdown
	res["markdown"] = h.markdown
	return res
}

// Markdown converts markdown to HTML replacing ONS custom tags
func (h *Helper) markdown(md string) template.HTML {
	s := h.replaceCustomTags(md)
	return helper.Markdown(s)
}

func (h *Helper) replaceCustomTags(text string) string {
	// Concurrently resolve figure data coming from zebedee
	var wg sync.WaitGroup
	// We use this buffered channel to limit the number of concurrent calls we make to zebedee
	sem := make(chan int, 10)

	for _, item := range h.contentResolvers {

		matches := item.Regexp.FindAllStringSubmatch(text, -1)

		for _, match := range matches {
			sem <- 1
			wg.Add(1)
			go func(rc RenderContent, match []string) {
				defer func() {
					<-sem
					wg.Done()
				}()

				partial, _ := rc(match)
				// TODO check err

				text = strings.Replace(text, match[0], partial, 1)
			}(item.RenderContent, match)
		}
	}

	wg.Wait()
	return text
}
