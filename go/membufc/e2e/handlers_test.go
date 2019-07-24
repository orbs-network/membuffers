// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package e2e

import (
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options/handlers"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options"
	"testing"
)

type implementsExampleService struct{}

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

type implementsServicesIProvideToOthersHandler struct{}

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
