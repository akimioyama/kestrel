package requests

import (
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/topics"
)

// CreateTopic — команда клиента создать топик с именем TopicName.
type CreateTopic struct {
	CorrelationID messages.CorrelationID
	TopicName     topics.TopicName
}

// Возвращает тип сообщения, соответствующий команде CreateTopic.
func (CreateTopic) Type() messages.MessageType {
	return messages.TypeCreateTopic
}
