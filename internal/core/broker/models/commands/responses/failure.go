package responses

import "kestrel/internal/core/broker/models/commands"

// Failure — неудачный ответ: команда не выполнена по ожидаемой причине
// (например, топик уже существует).
type Failure struct {
	CorrelationID commands.CorrelationID
	Code          ErrorCode
	Message       string
}

// NewFailure создаёт неудачный ответ.
func NewFailure(correlationID commands.CorrelationID, code ErrorCode, message string) Failure {
	return Failure{
		CorrelationID: correlationID,
		Code:          code,
		Message:       message,
	}
}

func (Failure) isResponse() {}
