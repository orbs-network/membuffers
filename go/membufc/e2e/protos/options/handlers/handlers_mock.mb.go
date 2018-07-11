// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package handlers

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service ServicesINeedFromOthersHandler

type MockServicesINeedFromOthersHandler struct {
	mock.Mock
}

func (s *MockServicesINeedFromOthersHandler) SomeMethodINeedFromOthers(input *SomeMessage) (*SomeMessage, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SomeMessage), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

/////////////////////////////////////////////////////////////////////////////
// service ServicesIProvideToOthersHandler

type MockServicesIProvideToOthersHandler struct {
	mock.Mock
}

func (s *MockServicesIProvideToOthersHandler) SomeMethodIProvideToOthers(input *SomeMessage) (*SomeMessage, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*SomeMessage), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

