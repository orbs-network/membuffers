// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package dep11

import (
)

/////////////////////////////////////////////////////////////////////////////
// enums

type DependencyEnum uint16

const (
	DEPENDENCY_ENUM_OPTION_A DependencyEnum = 0
	DEPENDENCY_ENUM_OPTION_B DependencyEnum = 1
)

func (n DependencyEnum) String() string {
	switch n {
	case DEPENDENCY_ENUM_OPTION_A:
		return "DEPENDENCY_ENUM_OPTION_A"
	case DEPENDENCY_ENUM_OPTION_B:
		return "DEPENDENCY_ENUM_OPTION_B"
	}
	return "UNKNOWN"
}

