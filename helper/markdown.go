package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"

	// render "github.com/ONSdigital/dp-renderer"
	// "github.com/ONSdigital/dp-renderer/assets"

	"github.com/ONSdigital/dp-renderer/assets"
	"github.com/russross/blackfriday/v2"
	"github.com/unrolled/render"
	unrolled "github.com/unrolled/render"
)

// Markdown converts markdown to HTML
func Markdown(md string) template.HTML {
	// lot's of the markdown we currently have stored doesn't match markdown title specs
	// currently it has no space between the hashes and the title text e.g. ##Title
	// to use our new markdown parser we have add a space e.g. ## Title
	re := regexp.MustCompile(`(##+)([^\s#])`)

	modifiedMarkdown := strings.Builder{}
	for _, line := range strings.Split(md, "\n") {
		modifiedMarkdown.WriteString(fmt.Sprintf("%s\n", re.ReplaceAllString(line, "$1 $2")))
	}

	s := blackfriday.Run([]byte(modifiedMarkdown.String()))
	return template.HTML(replaceCustomTags(string(s)))
}

func replaceCustomTags(md string) string {
	md = replaceChartTag(md)
	return md
}

type model struct{}

func replaceChartTag(md string) string {
	re := regexp.MustCompile(`<ons-chart path="(.*)" />`)

	isDevelopment := false

	model := model{} // Create model from path ($1)

	// rc := render.NewWithDefaultClient(assets.Asset, assets.AssetNames, "path", "siteDomain")
	buf := new(bytes.Buffer)

	unrolled.New(render.Options{
		Asset:         assets.Asset,
		AssetNames:    assets.AssetNames,
		Layout:        "",
		IsDevelopment: isDevelopment,
		Funcs:         []template.FuncMap{},
		// Funcs: []template.FuncMap{RegisteredFuncs},
	}).HTML(buf, 200, "partials/chart", model)

	return re.ReplaceAllString(md, buf.String())
}
