package topics

// Topic представляет топик брокера сообщений.
type Topic struct {
	ID   TopicID
	Name TopicName
}

// NewTopic создаёт новый топик с заданным ID и именем.
func NewTopic(id TopicID, name TopicName) Topic {
	return Topic{
		ID:   id,
		Name: name,
	}
}
