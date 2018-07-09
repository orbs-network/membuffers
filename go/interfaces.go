package membuffers

type Message interface {
	IsValid() bool
	Raw() []byte
}

type Builder interface {
	Write(buf []byte) (err error)
	GetSize() Offset
	CalcRequiredSize() Offset
	Build() Message
}
