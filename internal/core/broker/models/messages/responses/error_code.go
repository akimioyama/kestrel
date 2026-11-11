package responses

// ErrorCode — причина неудачного ответа.
type ErrorCode uint8

const (
	CodeUnknownTopic     ErrorCode = 0x01
	CodeTopicExists      ErrorCode = 0x02
	CodeInvalidHeader    ErrorCode = 0x03
	CodeUnknownType      ErrorCode = 0x04
	CodeInternal         ErrorCode = 0x05
	CodeInvalidTopicName ErrorCode = 0x06
)
