// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package types

import (
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/dep1"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorageNS

type StateStorageNS interface {
	WriteKeyNS(input *WriteKeyInputNS) (*WriteKeyOutputNS, error)
	ReadKeyNS(input *ReadKeyInputNS) (*ReadKeyOutputNS, error)
}

/////////////////////////////////////////////////////////////////////////////
// message WriteKeyInputNS (non serializable)

type WriteKeyInputNS struct {
	Key string
	Value uint32
}

/////////////////////////////////////////////////////////////////////////////
// message WriteKeyOutputNS (non serializable)

type WriteKeyOutputNS struct {
	Key string
	Additional []string
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeyInputNS (non serializable)

type ReadKeyInputNS struct {
	Key string
	Transaction *Transaction
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeyOutputNS (non serializable)

type ReadKeyOutputNS struct {
	Value uint32
	Dep []*dep1.DependencyMessage
}

/////////////////////////////////////////////////////////////////////////////
// enums

