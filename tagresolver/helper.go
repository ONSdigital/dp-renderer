package tagresolver

import (
	"context"
	"html/template"
	"regexp"
	"strings"
	"sync"

	render "github.com/ONSdigital/dp-renderer"
	"github.com/ONSdigital/dp-renderer/client"
	"github.com/ONSdigital/dp-renderer/helper"
	"github.com/ONSdigital/log.go/v2/log"
)

type RenderContent func(match []string) (string, error)
type GetContent func(path string) (interface{}, error)

type contentResolver struct {
	Regexp        regexp.Regexp
	RenderContent RenderContent
}

type TagResolverHelper struct {
	contentResolvers []contentResolver
	getContent       GetContent
	render           *render.Render
}

func NewTagResolverHelper(asset func(name string) ([]byte, error), assetNames func() []string, patternLibraryAssetsPath string, siteDomain string, getContent GetContent) *TagResolverHelper {
	isDevelopment := false
	if strings.Contains(siteDomain, "localhost") {
		isDevelopment = true
	}
	helper := &TagResolverHelper{
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

	equationResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-equation\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
		RenderContent: helper.ONSEquationResolver,
	}

	helper.contentResolvers = []contentResolver{
		chartResolver,
		tableResolver,
		equationResolver,
	}

	return helper
}

func (h *TagResolverHelper) GetFuncMap() template.FuncMap {
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
func (h *TagResolverHelper) markdown(md string) template.HTML {
	s := h.replaceCustomTags(md)
	return helper.Markdown(s)
}

func (h *TagResolverHelper) replaceCustomTags(text string) string {
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

				partial, err := rc(match)
				if err != nil {
					ctx := context.Background()
					log.Error(ctx, "failed to render custom tag", err, log.Data{"tag": match[0]})
				} else {
					text = strings.Replace(text, match[0], partial, 1)
				}

			}(item.RenderContent, match)
		}
	}

	wg.Wait()
	return text
}
