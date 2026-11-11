package responses

import "kestrel/internal/core/broker/models/commands"

// CreateTopicOK — успешный ответ: топик создан.
type CreateTopicOK struct {
	CorrelationID commands.CorrelationID
}

// NewCreateTopicOK создаёт успешный ответ на команду CreateTopic.
func NewCreateTopicOK(correlationID commands.CorrelationID) CreateTopicOK {
	return CreateTopicOK{CorrelationID: correlationID}
}

func (CreateTopicOK) isResponse() {}
