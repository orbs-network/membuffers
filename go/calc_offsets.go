//+build !race

// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package membuffers

func (m *InternalMessage) lazyCalcOffsets() bool {
	return m._lazyCalcOffsets()
}
