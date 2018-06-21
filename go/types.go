package membuffers

type FieldType uint16

const (
	TypeMessage    		FieldType = 1
	TypeBytes					FieldType = 2
	TypeString				FieldType = 3
	TypeUnion					FieldType = 4
	TypeUint8					FieldType = 11
	TypeUint16				FieldType = 12
	TypeUint32				FieldType = 13
	TypeUint64				FieldType = 14
	TypeUint8Array		FieldType = 21
	TypeUint16Array		FieldType = 22
	TypeUint32Array		FieldType = 23
	TypeUint64Array		FieldType = 24
	TypeMessageArray 	FieldType = 31
	TypeBytesArray		FieldType = 32
	TypeStringArray		FieldType = 33
)

var FieldSizes = map[FieldType]Offset{
	TypeMessage: 4,
	TypeBytes: 4,
	TypeString: 4,
	TypeUnion: 2,
	TypeUint8: 1,
	TypeUint16: 2,
	TypeUint32: 4,
	TypeUint64: 8,
	TypeUint8Array: 4,
	TypeUint16Array: 4,
	TypeUint32Array: 4,
	TypeUint64Array: 4,
	TypeMessageArray: 4,
	TypeBytesArray: 4,
	TypeStringArray: 4,
}

var FieldAlignment = map[FieldType]Offset{
	TypeMessage: 4,
	TypeBytes: 4,
	TypeString: 4,
	TypeUnion: 2,
	TypeUint8: 1,
	TypeUint16: 2,
	TypeUint32: 4,
	TypeUint64: 4,
	TypeUint8Array: 4,
	TypeUint16Array: 4,
	TypeUint32Array: 4,
	TypeUint64Array: 4,
	TypeMessageArray: 4,
	TypeBytesArray: 4,
	TypeStringArray: 4,
}

var FieldDynamic = map[FieldType]bool{
	TypeMessage: true,
	TypeBytes: true,
	TypeString: true,
	TypeUnion: true,
	TypeUint8: false,
	TypeUint16: false,
	TypeUint32: false,
	TypeUint64: false,
	TypeUint8Array: true,
	TypeUint16Array: true,
	TypeUint32Array: true,
	TypeUint64Array: true,
	TypeMessageArray: true,
	TypeBytesArray: true,
	TypeStringArray: true,
}

var FieldDynamicContentAlignment = map[FieldType]Offset{
	TypeMessage: 4,
	TypeBytes: 1,
	TypeString: 1,
	TypeUnion: 0,
	TypeUint8: 0,
	TypeUint16: 0,
	TypeUint32: 0,
	TypeUint64: 0,
	TypeUint8Array: 1,
	TypeUint16Array: 2,
	TypeUint32Array: 4,
	TypeUint64Array: 4,
	TypeMessageArray: 4,
	TypeBytesArray: 4,
	TypeStringArray: 4,
}