package responses

import (
	"fmt"
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
func NewFailure(correlationID messages.CorrelationID, code ErrorCode, message string) (Failure, error) {

	if message == "" {
		return Failure{}, fmt.Errorf("message cannot be empty")
	}

	if strings.TrimSpace(message) == "" {
		return Failure{}, fmt.Errorf("message cannot be whitespace only")
	}

	return Failure{
			CorrelationID: correlationID,
			Code:          code,
			Message:       message,
		},
		nil
}

func (Failure) isResponse() {}
