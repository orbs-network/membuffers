package membuffers

type ErrInvalidField struct {}
func (e *ErrInvalidField) Error() string { return "invalid field" }

type ErrSizeMismatch struct {}
func (e *ErrSizeMismatch) Error() string { return "size mismatch" }

type ErrBufferOverrun struct {}
func (e *ErrBufferOverrun) Error() string { return "buffer overrun" }