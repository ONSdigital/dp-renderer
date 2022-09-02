package sixteenstagresolver

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

type renderContent func(match []string) (string, error)

type contentResolver struct {
	Regexp        regexp.Regexp
	RenderContent renderContent
}

type TagResolverHelper struct {
	contentResolvers []contentResolver
	resourceReader   resourceReader
	render           *render.Render
}

type TagResolverRenderConfig struct {
	Asset                    func(name string) ([]byte, error)
	AssetNames               func() []string
	PatternLibraryAssetsPath string
	SiteDomain               string
}

func NewTagResolverHelper(uri string, rr ResourceReader, cfg TagResolverRenderConfig) *TagResolverHelper {
	isDevelopment := false
	if strings.Contains(cfg.SiteDomain, "localhost") {
		isDevelopment = true
	}
	resourceReader := resourceReader{reader: rr, uri: uri}
	helper := &TagResolverHelper{
		resourceReader: resourceReader,
		render:         render.New(client.NewUnrolledAdapterWithLayout(cfg.Asset, cfg.AssetNames, isDevelopment, ""), cfg.PatternLibraryAssetsPath, cfg.SiteDomain),
	}

	boxResolver := contentResolver{
		Regexp:        *regexp.MustCompile(`(?s)<ons-box align="([a-zA-Z]*)">(.*?)</ons-box>`),
		RenderContent: helper.ONSBoxResolver,
	}

	// chartResolver := contentResolver{
	// 	Regexp:        *regexp.MustCompile("<ons-chart\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
	// 	RenderContent: helper.ONSChartResolver,
	// }

	equationResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-equation\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
		RenderContent: helper.ONSEquationResolver,
	}

	imageResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-image\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
		RenderContent: helper.ONSImageResolver,
	}

	quoteResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-quote\\scontent=\"(.*?)\"\\s?(?:\\s+attr=\"(.*?)\")?\\s*/>"),
		RenderContent: helper.ONSQuoteResolver,
	}

	tableResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-table\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
		RenderContent: helper.ONSTableResolver,
	}

	tablev2Resolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-table-v2\\spath=\"([-A-Za-z0-9+&@#/%?=~_|!:,.;()*$]+)\"?\\s?/>"),
		RenderContent: helper.ONSTableV2Resolver,
	}

	warningResolver := contentResolver{
		Regexp:        *regexp.MustCompile(`(?s)<ons-warning-box>(.*?)</ons-warning-box>`),
		RenderContent: helper.ONSWarningResolver,
	}

	helper.contentResolvers = []contentResolver{
		boxResolver,
		// chartResolver,
		equationResolver,
		imageResolver,
		quoteResolver,
		tableResolver,
		tablev2Resolver,
		warningResolver,
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
			go func(rc renderContent, match []string) {
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
