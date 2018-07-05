// AUTO GENERATED FILE (by membufc proto compiler)
package options

import (
	"github.com/maraino/go-mock"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ExampleService

type MockExampleService struct {
	mock.Mock
	handlers.MockServicesINeedFromOthersHandler
}

func (s *MockExampleService) ExampleMethod(input *ExampleMethodInput) (*ExampleMethodOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ExampleMethodOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockExampleService) RegisterServicesIProvideToOthersHandler(handler handlers.ServicesIProvideToOthersHandler) {
	s.Called(handler)
}

