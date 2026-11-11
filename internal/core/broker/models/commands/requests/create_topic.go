package requests

import (
	"kestrel/internal/core/broker/models/commands"
	"kestrel/internal/core/broker/models/topics"
)

// CreateTopic — команда клиента создать топик с именем TopicName.
type CreateTopic struct {
	CorrelationID commands.CorrelationID
	TopicName     topics.TopicName
}

// NewCreateTopic создаёт команду CreateTopic.
func NewCreateTopic(correlationID commands.CorrelationID, topicName topics.TopicName) CreateTopic {
	return CreateTopic{
		CorrelationID: correlationID,
		TopicName:     topicName,
	}
}

func (CreateTopic) Type() RequestType {
	return TypeCreateTopic
}
