package onshelper

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"sync"

	"github.com/russross/blackfriday/v2"
)

type RenderContent func(match []string) (string, error)
type GetContent func(path string) (interface{}, error)

type contentResolver struct {
	Regexp        regexp.Regexp
	RenderContent RenderContent
}

type Helper struct {
	asset                    func(name string) ([]byte, error)
	assetNames               func() []string
	development              bool
	patternLibraryAssetsPath string
	siteDomain               string
	contentResolvers         []contentResolver
	getContent               GetContent
}

func NewHelper(asset func(name string) ([]byte, error), assetNames func() []string, patternLibraryAssetsPath string, siteDomain string, getContent GetContent) *Helper {
	isDevelopment := false
	if strings.Contains(siteDomain, "localhost") {
		isDevelopment = true
	}
	helper := &Helper{
		asset:                    asset,
		assetNames:               assetNames,
		development:              isDevelopment,
		patternLibraryAssetsPath: patternLibraryAssetsPath,
		siteDomain:               siteDomain,
		getContent:               getContent,
	}

	chartResolver := contentResolver{
		Regexp:        *regexp.MustCompile("<ons-chart path=\"(.*)\" />"),
		RenderContent: helper.ONSChartResolver,
	}

	helper.contentResolvers = []contentResolver{chartResolver}

	return helper
}

func (h *Helper) GetFuncMap() template.FuncMap {
	res := make(template.FuncMap)

	res["markdown"] = h.markdown

	return res
}

// Markdown converts markdown to HTML
func (h *Helper) markdown(md string) (template.HTML, error) {
	// lot's of the markdown we currently have stored doesn't match markdown title specs
	// currently it has no space between the hashes and the title text e.g. ##Title
	// to use our new markdown parser we have add a space e.g. ## Title
	re := regexp.MustCompile(`(##+)([^\s#])`)

	modifiedMarkdown := strings.Builder{}
	for _, line := range strings.Split(md, "\n") {
		modifiedMarkdown.WriteString(fmt.Sprintf("%s\n", re.ReplaceAllString(line, "$1 $2")))
	}

	s := string(blackfriday.Run([]byte(modifiedMarkdown.String())))

	s = h.replaceCustomTags(s)

	output := template.HTML(s)

	return output, nil
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
			go func(tag string) {
				defer func() {
					<-sem
					wg.Done()
				}()

				partial, _ := item.RenderContent(match)
				// TODO check err

				text = strings.Replace(text, tag, partial, 1)
			}(match[0])
		}
	}

	wg.Wait()
	return text
}

// func replaceCustomTags(md string) string {
// 	md = replaceChartTag(md)
// 	return md
// }

// func replaceChartTag(md string) string {
// 	re := regexp.MustCompile(`<ons-chart path="(.*)" />`)

// 	isDevelopment := false

// 	model := model{} // Create model from path ($1)

// 	// rc := render.NewWithDefaultClient(assets.Asset, assets.AssetNames, "path", "siteDomain")
// 	buf := new(bytes.Buffer)

// 	unrolled.New(render.Options{
// 		Asset:         assets.Asset,
// 		AssetNames:    assets.AssetNames,
// 		Layout:        "",
// 		IsDevelopment: isDevelopment,
// 		Funcs:         []template.FuncMap{},
// 		// Funcs: []template.FuncMap{RegisteredFuncs},
// 	}).HTML(buf, 200, "partials/chart", model)

// 	return re.ReplaceAllString(md, buf.String())
// }
