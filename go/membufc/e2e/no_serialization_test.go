package e2e

import (
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"testing"
	"errors"
)

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

func TestNSContainers(t *testing.T) {
	// build
	container := &types.ExampleContainer{
		Message1: (&types.MessageInContainerBuilder{
			Field: "hello",
		}).Build(),
		Container1: &types.NestedContainer{
			Name: "john",
		},
		Containers2: []*types.NestedContainer{
			{Name: "lucy"},
			{Name: "linda"},
			{Name: "nancy"},
		},
	}

	// read
	if container.Message1.Field() != "hello" {
		t.Fatalf("Message1: instead of expected got %v", container.Message1.Field())
	}
	if container.Container1.Name != "john" {
		t.Fatalf("Container1: instead of expected got %v", container.Container1.Name)
	}
	if container.Containers2[0].Name != "lucy" {
		t.Fatalf("Containers2[0]: instead of expected got %v", container.Containers2[0].Name)
	}
	if container.Containers2[1].Name != "linda" {
		t.Fatalf("Containers2[1]: instead of expected got %v", container.Containers2[0].Name)
	}
	if len(container.Containers2) != 3 {
		t.Fatalf("Containers2: instead of expected len got %v", len(container.Containers2))
	}
	if container.String() != `{Message1:{Field:hello,},Container1:{Name:john,},Containers2:[{Name:lucy,},{Name:linda,},{Name:nancy,},],}` {
		t.Fatalf("String: instead of expected got %s", container.String())
	}
}

func TestEmptyNSContainer(t *testing.T) {
	emptyContainer := &types.ExampleContainer{}

	if emptyContainer.String() != `{Message1:<nil>,Container1:<nil>,Containers2:[],}` {
		t.Fatalf("empty container String returned %s", emptyContainer.String())
	}
}