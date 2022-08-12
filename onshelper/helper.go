package onshelper

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"

	"github.com/russross/blackfriday/v2"
)

type RenderContent func(getContent GetContent, match []string) (string, error)
type GetContent func(path string) ([]byte, error)

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

func NewHelper(asset func(name string) ([]byte, error), assetNames func() []string, development bool, patternLibraryAssetsPath string, siteDomain string, getContent GetContent) *Helper {

	helper := &Helper{
		asset:                    asset,
		assetNames:               assetNames,
		development:              development,
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
	for _, item := range h.contentResolvers {

		matches := item.Regexp.FindAllStringSubmatch(s, -1)

		for _, match := range matches {
			partial, err := item.RenderContent(nil, match)

			if err != nil {
				return "", err
			}

			s = strings.Replace(s, match[0], partial, 1)
		}
	}

	output := template.HTML(s)

	return output, nil
}
