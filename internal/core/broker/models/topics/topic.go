package topics

type Topic struct {
	InternalID TopicID
	Name       TopicName
}

func NewTopic(internalID TopicID, name TopicName) Topic {
	return Topic{
		InternalID: internalID,
		Name:       name,
	}
}
