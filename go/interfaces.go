package membuffers

type Message interface {
	IsValid() bool
	Raw() []byte
	String() string
}

type Builder interface {
	Write(buf []byte) (err error)
	GetSize() Offset
	CalcRequiredSize() Offset
	Build() Message
}
