package e2e

import (
	"context"
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"github.com/pkg/errors"
	"testing"
)

type stateStorageWithCtxService struct {
}

func (s *stateStorageWithCtxService) WriteKeyWithCtx(context.Context, *types.WriteKeyWithCtxInput) (*types.WriteKeyWithCtxOutput, error) {
	return nil, nil
}

func (s *stateStorageWithCtxService) ReadKeyWithCtx(context.Context, *types.ReadKeyWithCtxInput) (*types.ReadKeyWithCtxOutput, error) {
	return nil, nil
}

func NewStateStorageWithCtx() types.StateStorageWithCtx {
	return &stateStorageWithCtxService{}
}

func TestServiceWithCtxMock(t *testing.T) {
	m := &types.MockStateStorageWithCtx{}

	wantedWriteOut := &types.WriteKeyWithCtxOutput{}
	m.When("WriteKeyWithCtx", mock.Any, nil).Return(wantedWriteOut, nil).Times(1)
	writeOut, err := m.WriteKeyWithCtx(context.TODO(), nil)
	if writeOut != wantedWriteOut {
		t.Fatalf("Mock WriteKeyWithCtxOutput is not as expected")
	}
	if err != nil {
		t.Fatalf("Mock err is not nil")
	}

	m.When("WriteKeyWithCtx", mock.Any, nil).Return(nil, errors.New("errorWrite")).Times(1)
	writeOut, err = m.WriteKeyWithCtx(context.TODO(), nil)
	if writeOut != nil {
		t.Fatalf("Mock WriteKeyWithCtxOutput is not nil")
	}
	if err.Error() != "errorWrite" {
		t.Fatalf("Mock err is not as expected")
	}

	wantedReadOut := &types.ReadKeyWithCtxOutput{}
	m.When("ReadKeyWithCtx", mock.Any, nil).Return(wantedReadOut, nil).Times(1)
	readOut, err := m.ReadKeyWithCtx(context.TODO(), nil)
	if readOut != wantedReadOut {
		t.Fatalf("Mock ReadKeyWithCtxOutput is not as expected")
	}
	if err != nil {
		t.Fatalf("Mock err is not nil")
	}

	m.When("ReadKeyWithCtx", mock.Any, nil).Return(nil, errors.New("errorRead")).Times(1)
	readOut, err = m.ReadKeyWithCtx(context.TODO(), nil)
	if readOut != nil {
		t.Fatalf("Mock ReadKeyWithCtxOutput is not nil")
	}
	if err.Error() != "errorRead" {
		t.Fatalf("Mock err is not as expected")
	}
}
