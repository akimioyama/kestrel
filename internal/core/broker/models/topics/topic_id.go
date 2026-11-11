package topics

type TopicID struct {
	Value uint64
}

func NewTopicID(value uint64) TopicID {
	return TopicID{Value: value}
}
