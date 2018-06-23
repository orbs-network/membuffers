package membuffers

import (
	"math"
	"unsafe"
)

type (
	Offset uint32
)

func GetByte(buf []byte) byte {
	return byte(GetUint8(buf))
}

func GetUint8(buf []byte) (n uint8) {
	return *(*uint8)(unsafe.Pointer(&buf[0]))
}

func GetUint8Polyfill(buf []byte) (n uint8) {
	n = uint8(buf[0])
	return
}

func GetUint16(buf []byte) uint16 {
	return *(*uint16)(unsafe.Pointer(&buf[0]))
}

func GetUint16Polyfill(buf []byte) (n uint16) {
	n |= uint16(buf[0])
	n |= uint16(buf[1]) << 8
	return
}

func GetUint32(buf []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&buf[0]))
}

func GetUint32Polyfill(buf []byte) (n uint32) {
	n |= uint32(buf[0])
	n |= uint32(buf[1]) << 8
	n |= uint32(buf[2]) << 16
	n |= uint32(buf[3]) << 24
	return
}

func GetUint64(buf []byte) (n uint64) {
	n |= uint64(*(*uint32)(unsafe.Pointer(&buf[0])))
	n |= uint64(*(*uint32)(unsafe.Pointer(&buf[4]))) << 32
	return
}

func GetUint64Polyfill(buf []byte) (n uint64) {
	n |= uint64(buf[0])
	n |= uint64(buf[1]) << 8
	n |= uint64(buf[2]) << 16
	n |= uint64(buf[3]) << 24
	n |= uint64(buf[4]) << 32
	n |= uint64(buf[5]) << 40
	n |= uint64(buf[6]) << 48
	n |= uint64(buf[7]) << 56
	return
}

func GetInt8(buf []byte) int8 {
	return *(*int8)(unsafe.Pointer(&buf[0]))
}

func GetInt8Polyfill(buf []byte) (n int8) {
	n = int8(buf[0])
	return
}

func GetInt16(buf []byte) int16 {
	return *(*int16)(unsafe.Pointer(&buf[0]))
}

func GetInt16Polyfill(buf []byte) (n int16) {
	n |= int16(buf[0])
	n |= int16(buf[1]) << 8
	return
}

func GetInt32(buf []byte) int32 {
	return *(*int32)(unsafe.Pointer(&buf[0]))
}

func GetInt32Polyfill(buf []byte) (n int32) {
	n |= int32(buf[0])
	n |= int32(buf[1]) << 8
	n |= int32(buf[2]) << 16
	n |= int32(buf[3]) << 24
	return
}

func GetInt64(buf []byte) int64 {
	return *(*int64)(unsafe.Pointer(&buf[0]))
}

func GetInt64Polyfill(buf []byte) (n int64) {
	n |= int64(buf[0])
	n |= int64(buf[1]) << 8
	n |= int64(buf[2]) << 16
	n |= int64(buf[3]) << 24
	n |= int64(buf[4]) << 32
	n |= int64(buf[5]) << 40
	n |= int64(buf[6]) << 48
	n |= int64(buf[7]) << 56
	return
}

func GetFloat32(buf []byte) float32 {
	x := GetUint32(buf)
	return math.Float32frombits(x)
}

func GetFloat64(buf []byte) float64 {
	x := GetUint64(buf)
	return math.Float64frombits(x)
}

func GetOffset(buf []byte) Offset {
	return Offset(GetUint32(buf))
}

func GetUnionType(buf []byte) uint16 {
	return GetUint16(buf)
}

func WriteByte(buf []byte, n byte) {
	WriteUint8(buf, uint8(n))
}

func WriteUint8(buf []byte, n uint8) {
	*(*uint8)(unsafe.Pointer(&buf[0])) = n
}

func WriteUint8Polyfill(buf []byte, n uint8) {
	buf[0] = byte(n)
}

func WriteUint16(buf []byte, n uint16) {
	*(*uint16)(unsafe.Pointer(&buf[0])) = n
}

func WriteUint16Polyfill(buf []byte, n uint16) {
	buf[0] = byte(n)
	buf[1] = byte(n >> 8)
}

func WriteUint32(buf []byte, n uint32) {
	*(*uint32)(unsafe.Pointer(&buf[0])) = n
}

func WriteUint32Polyfill(buf []byte, n uint32) {
	buf[0] = byte(n)
	buf[1] = byte(n >> 8)
	buf[2] = byte(n >> 16)
	buf[3] = byte(n >> 24)
}

func WriteUint64(buf []byte, n uint64) {
	*(*uint32)(unsafe.Pointer(&buf[0])) = uint32(n)
	*(*uint32)(unsafe.Pointer(&buf[4])) = uint32(n >> 32)
}

func WriteUint64Polyfill(buf []byte, n uint64) {
	for i := uint(0); i < uint(8); i++ {
		buf[i] = byte(n >> (i * 8))
	}
}

func WriteInt8(buf []byte, n int8) {
	*(*int8)(unsafe.Pointer(&buf[0])) = n
}

func WriteInt8Polyfill(buf []byte, n int8) {
	buf[0] = byte(n)
}

func WriteInt16(buf []byte, n int16) {
	*(*int16)(unsafe.Pointer(&buf[0])) = n
}

func WriteInt16Polyfill(buf []byte, n int16) {
	buf[0] = byte(n)
	buf[1] = byte(n >> 8)
}

func WriteInt32(buf []byte, n int32) {
	*(*int32)(unsafe.Pointer(&buf[0])) = n
}

func WriteInt32Polyfill(buf []byte, n int32) {
	buf[0] = byte(n)
	buf[1] = byte(n >> 8)
	buf[2] = byte(n >> 16)
	buf[3] = byte(n >> 24)
}

func WriteInt64(buf []byte, n int64) {
	*(*uint32)(unsafe.Pointer(&buf[0])) = uint32(n)
	*(*uint32)(unsafe.Pointer(&buf[4])) = uint32(n >> 32)
}

func WriteInt64Polyfill(buf []byte, n int64) {
	for i := uint(0); i < uint(8); i++ {
		buf[i] = byte(n >> (i * 8))
	}
}

func WriteFloat32(buf []byte, n float32) {
	WriteUint32(buf, math.Float32bits(n))
}

func WriteFloat64(buf []byte, n float64) {
	WriteUint64(buf, math.Float64bits(n))
}

func WriteOffset(buf []byte, n Offset) {
	WriteUint32(buf, uint32(n))
}

func WriteUnionType(buf []byte, n uint16) {
	WriteUint16(buf, n)
}

func byteSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}