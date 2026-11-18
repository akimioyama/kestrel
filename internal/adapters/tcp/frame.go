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

// FieldID полей Header.
const (
	fieldTopic         uint8 = 0x01
	fieldCorrelationID uint8 = 0x02
	fieldKey           uint8 = 0x03
	fieldErrorCode     uint8 = 0x04
	fieldErrorMessage  uint8 = 0x05
	fieldMessageID     uint8 = 0x06
)
