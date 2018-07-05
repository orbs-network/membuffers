// AUTO GENERATED FILE (by membufc proto compiler)
package options

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ExampleService

type MockExampleService struct {
	mock.Mock
}

func (s *MockExampleService) ExampleMethod(input *ExampleMethodInput) (*ExampleMethodOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ExampleMethodOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

