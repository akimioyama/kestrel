package responses

import "kestrel/internal/core/broker/models/messages"

// CreateTopicOK — успешный ответ: топик создан.
type CreateTopicOK struct {
	CorrelationID messages.CorrelationID
}

func (CreateTopicOK) isResponse() {}
