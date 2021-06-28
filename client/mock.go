package client

import (
	"errors"
	"fmt"
	"io"

	"github.com/ONSdigital/dp-renderer/model"
)

type MockRenderClient struct {
	TemplateNames             []string
	ValidBuildHTMLMethodCalls int
	ValidSetErrorMethodCalls  int
}

func NewMockRenderingClient(templateNames []string) *MockRenderClient {
	return &MockRenderClient{
		TemplateNames: templateNames,
	}
}

func (m *MockRenderClient) BuildHTML(w io.Writer, status int, templateName string, pageModel interface{}) error {
	for _, value := range m.TemplateNames {
		if value == templateName {
			fmt.Println(value)
			m.ValidBuildHTMLMethodCalls++
			return nil
		}
	}
	return errors.New("Failed to build page")
}

func (m *MockRenderClient) SetError(w io.Writer, status int, errorModel model.ErrorResponse) error {
	m.ValidSetErrorMethodCalls++
	return errors.New("An error occurred when building the page")
}
