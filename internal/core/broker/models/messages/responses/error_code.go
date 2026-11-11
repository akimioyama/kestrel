package responses

// ErrorCode — причина неудачного ответа.
type ErrorCode uint8

const (
	// CodeUnknownTopic — указанный топик не найден.
	CodeUnknownTopic ErrorCode = 0x01

	// CodeTopicExists — топик с таким именем уже существует.
	CodeTopicExists ErrorCode = 0x02

	// CodeInvalidHeader — некорректный заголовок сообщения.
	CodeInvalidHeader ErrorCode = 0x03

	// CodeUnknownType — неизвестный тип сообщения.
	CodeUnknownType ErrorCode = 0x04

	// CodeInternal — внутренняя ошибка сервера.
	CodeInternal ErrorCode = 0x05

	// CodeInvalidTopicName — недопустимое имя топика.
	CodeInvalidTopicName ErrorCode = 0x06
)
