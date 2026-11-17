package responses

import (
	"kestrel/internal/core/broker/models/messages"
	"strings"
)

// Failure — неудачный ответ: команда не выполнена по ожидаемой причине
// (например, топик уже существует).
type Failure struct {
	CorrelationID messages.CorrelationID
	Code          ErrorCode
	Message       string
}

// NewFailure создаёт неудачный ответ.
// Паникует, если message пустое или состоит только из пробелов:
// сообщение всегда задаётся в коде, поэтому пустое значение — ошибка программиста.
func NewFailure(correlationID messages.CorrelationID, code ErrorCode, message string) Failure {
	if strings.TrimSpace(message) == "" {
		panic("responses.NewFailure: message cannot be empty or whitespace only")
	}

	return Failure{
		CorrelationID: correlationID,
		Code:          code,
		Message:       message,
	}
}

func (Failure) isResponse() {}
