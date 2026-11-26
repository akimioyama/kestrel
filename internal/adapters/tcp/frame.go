package tcp

// Frame — один фрейм протокола: тип операции, поля Header и Payload.
type Frame struct {
	Type    uint8
	Header  Header
	Payload []byte
}

// Header — поля из Header фрейма.
type Header struct {
	Topic         *string
	CorrelationID *uint64
	Key           []byte
	ErrorCode     *uint8
	ErrorMessage  *string
	MessageID     *uint64
}
