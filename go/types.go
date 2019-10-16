// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package membuffers

type FieldType uint16

const (
	TypeMessage      FieldType = 1
	TypeBytes        FieldType = 2
	TypeString       FieldType = 3
	TypeUnion        FieldType = 4
	TypeBool         FieldType = 10
	TypeUint8        FieldType = 11
	TypeUint16       FieldType = 12
	TypeUint32       FieldType = 13
	TypeUint64       FieldType = 14
	TypeUint256      FieldType = 16
	TypeUint8Array   FieldType = 21
	TypeUint16Array  FieldType = 22
	TypeUint32Array  FieldType = 23
	TypeUint64Array  FieldType = 24
	TypeUint256Array FieldType = 26
	TypeMessageArray FieldType = 31
	TypeBytesArray   FieldType = 32
	TypeStringArray  FieldType = 33

	TypeBytes32      FieldType = 41
	TypeBytes20      FieldType = 42
	TypeBytes32Array FieldType = 51
	TypeBytes20Array FieldType = 52
)

var FieldSizes = map[FieldType]Offset{
	TypeMessage:      4,
	TypeBytes:        4,
	TypeString:       4,
	TypeUnion:        2,
	TypeBool:         1,
	TypeUint8:        1,
	TypeUint16:       2,
	TypeUint32:       4,
	TypeUint64:       8,
	TypeUint256:      32,
	TypeUint8Array:   4,
	TypeUint16Array:  4,
	TypeUint32Array:  4,
	TypeUint64Array:  4,
	TypeUint256Array: 4,
	TypeMessageArray: 4,
	TypeBytesArray:   4,
	TypeStringArray:  4,

	TypeBytes32:      32,
	TypeBytes20:      20,
	TypeBytes32Array: 4,
	TypeBytes20Array: 4,
}

var FieldAlignment = map[FieldType]Offset{
	TypeMessage:      4,
	TypeBytes:        4,
	TypeString:       4,
	TypeUnion:        2,
	TypeBool:         1,
	TypeUint8:        1,
	TypeUint16:       2,
	TypeUint32:       4,
	TypeUint64:       4,
	TypeUint256:      4,
	TypeUint8Array:   4,
	TypeUint16Array:  4,
	TypeUint32Array:  4,
	TypeUint64Array:  4,
	TypeUint256Array: 4,
	TypeMessageArray: 4,
	TypeBytesArray:   4,
	TypeStringArray:  4,

	TypeBytes32:      4,
	TypeBytes20:      4,
	TypeBytes32Array: 4,
	TypeBytes20Array: 4,
}

var FieldDynamic = map[FieldType]bool{
	TypeMessage:      true,
	TypeBytes:        true,
	TypeString:       true,
	TypeUnion:        true,
	TypeBool:         false,
	TypeUint8:        false,
	TypeUint16:       false,
	TypeUint32:       false,
	TypeUint64:       false,
	TypeUint256:      false,
	TypeUint8Array:   true,
	TypeUint16Array:  true,
	TypeUint32Array:  true,
	TypeUint64Array:  true,
	TypeUint256Array: true,
	TypeMessageArray: true,
	TypeBytesArray:   true,
	TypeStringArray:  true,

	TypeBytes32:      false,
	TypeBytes20:      false,
	TypeBytes32Array: true,
	TypeBytes20Array: true,
}

var FieldDynamicContentAlignment = map[FieldType]Offset{
	TypeMessage:      4,
	TypeBytes:        1,
	TypeString:       1,
	TypeUnion:        0,
	TypeBool:         0,
	TypeUint8:        0,
	TypeUint16:       0,
	TypeUint32:       0,
	TypeUint64:       0,
	TypeUint256:      0,
	TypeUint8Array:   1,
	TypeUint16Array:  2,
	TypeUint32Array:  4,
	TypeUint64Array:  4,
	TypeUint256Array: 4,
	TypeMessageArray: 4,
	TypeBytesArray:   4,
	TypeStringArray:  4,

	TypeBytes32:      0,
	TypeBytes20:      0,
	TypeBytes32Array: 4,
	TypeBytes20Array: 4,
}
