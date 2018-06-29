// AUTO GENERATED FILE (by membufc proto compiler)
package types

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type MockStateStorage struct {
	mock.Mock
}

func (s *MockStateStorage) WriteKey(input *WriteKeyInput) (*WriteKeyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*WriteKeyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorage) ReadKey(input *ReadKeyInput) (*ReadKeyOutput, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ReadKeyOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

