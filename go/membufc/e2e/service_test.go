package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"github.com/pkg/errors"
)

type stateStorageService struct {
	types.StateStorage
}

func (s *stateStorageService) WriteKey(*types.WriteKeyInput) (*types.WriteKeyOutput, error) {
	return nil, nil
}

func (s *stateStorageService) ReadKey(*types.ReadKeyInput) (*types.ReadKeyOutput, error) {
	return nil, nil
}

func TestServiceMock(t *testing.T) {
	m := &types.MockStateStorage{}

	wantedWriteOut := &types.WriteKeyOutput{}
	m.When("WriteKey", nil).Return(wantedWriteOut, nil).Times(1)
	writeOut, err := m.WriteKey(nil)
	if writeOut != wantedWriteOut {
		t.Fatalf("Mock WriteKeyOutput is not as expected")
	}
	if err != nil {
		t.Fatalf("Mock err is not nil")
	}

	m.When("WriteKey", nil).Return(nil, errors.New("errorWrite")).Times(1)
	writeOut, err = m.WriteKey(nil)
	if writeOut != nil {
		t.Fatalf("Mock WriteKeyOutput is not nil")
	}
	if err.Error() != "errorWrite" {
		t.Fatalf("Mock err is not as expected")
	}

	wantedReadOut := &types.ReadKeyOutput{}
	m.When("ReadKey", nil).Return(wantedReadOut, nil).Times(1)
	readOut, err := m.ReadKey(nil)
	if readOut != wantedReadOut {
		t.Fatalf("Mock ReadKeyOutput is not as expected")
	}
	if err != nil {
		t.Fatalf("Mock err is not nil")
	}

	m.When("ReadKey", nil).Return(nil, errors.New("errorRead")).Times(1)
	readOut, err = m.ReadKey(nil)
	if readOut != nil {
		t.Fatalf("Mock ReadKeyOutput is not nil")
	}
	if err.Error() != "errorRead" {
		t.Fatalf("Mock err is not as expected")
	}
}

type stateStorageServiceNS struct {
	types.StateStorageNS
}

func (s *stateStorageServiceNS) WriteKeyNS(*types.WriteKeyInputNS) (*types.WriteKeyOutputNS, error) {
	return nil, nil
}

func (s *stateStorageServiceNS) ReadKeyNS(*types.ReadKeyInputNS) (*types.ReadKeyOutputNS, error) {
	return nil, nil
}

func TestServiceNSMock(t *testing.T) {
	m := &types.MockStateStorageNS{}

	wantedWriteOut := &types.WriteKeyOutputNS{}
	m.When("WriteKeyNS", nil).Return(wantedWriteOut, nil).Times(1)
	writeOut, err := m.WriteKeyNS(nil)
	if writeOut != wantedWriteOut {
		t.Fatalf("Mock WriteKeyOutputNS is not as expected")
	}
	if err != nil {
		t.Fatalf("Mock err is not nil")
	}

	m.When("WriteKeyNS", nil).Return(nil, errors.New("errorWrite")).Times(1)
	writeOut, err = m.WriteKeyNS(nil)
	if writeOut != nil {
		t.Fatalf("Mock WriteKeyOutputNS is not nil")
	}
	if err.Error() != "errorWrite" {
		t.Fatalf("Mock err is not as expected")
	}

	wantedReadOut := &types.ReadKeyOutputNS{}
	m.When("ReadKeyNS", nil).Return(wantedReadOut, nil).Times(1)
	readOut, err := m.ReadKeyNS(nil)
	if readOut != wantedReadOut {
		t.Fatalf("Mock ReadKeyOutputNS is not as expected")
	}
	if err != nil {
		t.Fatalf("Mock err is not nil")
	}

	m.When("ReadKeyNS", nil).Return(nil, errors.New("errorRead")).Times(1)
	readOut, err = m.ReadKeyNS(nil)
	if readOut != nil {
		t.Fatalf("Mock ReadKeyOutputNS is not nil")
	}
	if err.Error() != "errorRead" {
		t.Fatalf("Mock err is not as expected")
	}
}
