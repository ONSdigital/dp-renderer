package tagresolver

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/ONSdigital/dp-renderer/helper"
	"github.com/ONSdigital/dp-renderer/model"
	"github.com/ONSdigital/log.go/v2/log"
	uuid "github.com/satori/go.uuid"
)

func (h *TagResolverHelper) ONSBoxResolver(language string) func([]string) (string, error) {
	return func(match []string) (string, error) {
		model := model.Figure{}
		model.Language = language

		// figureTag := match[0]   // figure tag
		model.Align = match[1]   // align attribute
		model.Content = match[2] // tag content
		return h.applyTemplate(model, "partials/ons-tags/ons-box"), nil
	}
}

func (h *TagResolverHelper) ONSChartResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path
	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}
	figure.DownloadFormats = []string{"csv", "xls"}
	return h.applyTemplate(figure, "partials/sixteens-ons-tags/ons-chart"), nil
}

func (h *TagResolverHelper) ONSEquationResolver(match []string) (string, error) {
	// figureTag := match[0]   // figure tag
	contentPath := match[1] // figure path

	figure, err := h.resourceReader.GetFigure(contentPath)
	if err != nil {
		return "", err
	}

	for i, sidecarFile := range figure.Files {
		if sidecarFile.Type == "generated-svg" {
			resource, err := h.resourceReader.GetResourceBody(sidecarFile.Filename)
			if err != nil {
				return "", err
			}
			figure.Files[i].Content = string(resource)
		}
	}

	return h.applyTemplate(figure, "partials/ons-tags/ons-equation"), nil
}

func (h *TagResolverHelper) ONSImageResolver(language string) func([]string) (string, error) {
	return func(match []string) (string, error) {
		// figureTag := match[0]   // figure tag
		contentPath := match[1] // figure path

		figure, err := h.resourceReader.GetFigure(contentPath)
		if err != nil {
			return "", err
		}

		figure.Language = language

		for i, sidecarFile := range figure.Files {
			size, err := h.resourceReader.GetFileSize(sidecarFile.Filename)
			if err != nil {
				log.Warn(context.Background(), "Image resolver couldn't find file size for image file using filename",
					log.Data{"file": h.resourceReader.getPathUri(sidecarFile.Filename), "path": contentPath})

				// If the file can't be found, try locating it by concatenating the file type to the path
				// This is a fallback for a legacy issue, mimicking what Babbage used to do
				filename := contentPath + "." + sidecarFile.FileType
				size, err = h.resourceReader.GetFileSize(filename)
				if err != nil {
					size = 0
					log.Error(context.Background(), "Image resolver couldn't find file size for image file using filetype",
						err, log.Data{"file": h.resourceReader.getPathUri(filename), "path": contentPath})
				}
			}
			figure.Files[i].FileSize = humanReadableByteCount(size)
		}
		return h.applyTemplate(figure, "partials/ons-tags/ons-image"), nil
	}
}

func (h *TagResolverHelper) ONSInteractiveResolver(language string) func([]string) (string, error) {
	return func(match []string) (string, error) {
		hasHypertextProtocol := func(s string) bool {
			return strings.HasPrefix(s, "http") || strings.HasPrefix(s, "https")
		}

		buildIframe := func(uri string) string {
			src := uri
			if !hasHypertextProtocol(uri) {
				src = "http://localhost" + uri
			}
			return fmt.Sprintf("<iframe width=\"100%%\" src=\"%s\"></iframe>", src)
		}

		model := model.Figure{}
		model.Language = language

		// figureTag := match[0]   // figure tag
		model.URI = match[1] // url
		model.Iframe = buildIframe(model.URI)

		if len(match) > 2 {
			model.FullWidth = match[2] == "true" // full-width
		}
		if len(match) > 3 {
			model.Title = match[3] // title
		}

		if model.Title == "" {
			model.Title = helper.Localise("OnsTagInteractiveChart", language, 1)
		}

		model.Id = uuid.NewV4().String()[:10]
		return h.applyTemplate(model, "partials/ons-tags/ons-interactive"), nil
	}
}

func (h *TagResolverHelper) ONSQuoteResolver(match []string) (string, error) {
	model := model.Figure{}
	// figureTag := match[0]   // figure tag
	model.Content = match[1] // content attribute
	if len(match) > 2 {
		model.Attribution = match[2] // attr attribute
	}

	return h.applyTemplate(model, "partials/ons-tags/ons-quote"), nil
}

func (h *TagResolverHelper) ONSTableResolver(language string) func([]string) (string, error) {
	return func(match []string) (string, error) {
		// figureTag := match[0]   // figure tag
		contentPath := match[1] // figure path
		figure, err := h.resourceReader.GetFigure(contentPath)
		if err != nil {
			return "", err
		}

		figure.Language = language

		for i, sidecarFile := range figure.Files {
			if sidecarFile.Type == "html" {
				resource, err := h.resourceReader.GetResourceBody(sidecarFile.Filename)
				if err != nil {
					return "", err
				}
				figure.Files[i].Content = string(resource)
			}

			size, err := h.resourceReader.GetFileSize(sidecarFile.Filename)
			if err != nil {
				return "", err
			}
			figure.Files[i].FileSize = humanReadableByteCount(size)
		}

		return h.applyTemplate(figure, "partials/ons-tags/ons-table"), nil
	}
}

func (h *TagResolverHelper) ONSTableV2Resolver(language string) func([]string) (string, error) {
	return func(match []string) (string, error) {
		// figureTag := match[0]   // figure tag
		contentPath := match[1] // figure path

		figureJSON, err := h.resourceReader.GetResourceBody(contentPath + ".json")
		if err != nil {
			return "", err
		}

		tableHTML, err := h.resourceReader.GetTable(figureJSON)
		if err != nil {
			return "", err
		}

		figure, err := h.resourceReader.GetFigure(contentPath)
		if err != nil {
			return "", err
		}
		figure.Language = language
		figure.Content = tableHTML

		return h.applyTemplate(figure, "partials/ons-tags/ons-table-v2"), nil
	}
}

func (h *TagResolverHelper) ONSWarningResolver(language string) func([]string) (string, error) {
	return func(match []string) (string, error) {
		model := model.Figure{}
		model.Language = language

		// figureTag := match[0]   // figure tag
		model.Content = match[1] // tag content

		return h.applyTemplate(model, "partials/ons-tags/ons-warning"), nil
	}
}

func (h *TagResolverHelper) applyTemplate(figure interface{}, template string) string {
	buf := new(bytes.Buffer)
	h.render.BuildPage(buf, figure, template)
	return buf.String()
}
