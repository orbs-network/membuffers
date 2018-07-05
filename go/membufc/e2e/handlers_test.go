package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options/handlers"
)

type implementsExampleService struct {}

func NewImplementsExampleService() options.ExampleService {
	return &implementsExampleService{}
}

func (m *implementsExampleService) ExampleMethod(input *options.ExampleMethodInput) (*options.ExampleMethodOutput, error) {
	return nil, nil
}

func (m *implementsExampleService) RegisterServicesIProvideToOthersHandler(handlers.ServicesIProvideToOthersHandler) {
}

func (m *implementsExampleService) SomeMethodINeedFromOthers(*handlers.SomeMessage) (*handlers.SomeMessage, error) {
	return nil, nil
}

type implementsServicesIProvideToOthersHandler struct {}

func NewImplementsServicesIProvideToOthersHandler() handlers.ServicesIProvideToOthersHandler {
	return &implementsServicesIProvideToOthersHandler{}
}

func (m *implementsServicesIProvideToOthersHandler) SomeMethodIProvideToOthers(*handlers.SomeMessage) (*handlers.SomeMessage, error) {
	return nil, nil
}

func TestHandlersInOptions(t *testing.T) {
	s := NewImplementsExampleService()
	out, err := s.ExampleMethod(nil)
	if out != nil || err != nil {
		t.Fatalf("ExampleMethod did not return as expected")
	}
	out2, err := s.SomeMethodINeedFromOthers(nil)
	if out2 != nil || err != nil {
		t.Fatalf("SomeMethodINeedFromOthers did not return as expected")
	}
	p := NewImplementsServicesIProvideToOthersHandler()
	s.RegisterServicesIProvideToOthersHandler(p)
}