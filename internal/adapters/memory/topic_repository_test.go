package memory

import (
	"kestrel/internal/core/broker/models/topics"
	"testing"
)

func TestTopicRepositoryCanAddTopic(test *testing.T) {
	repository := NewTopicRepository()

	topic := topics.Topic{
		ID:   topics.TopicID(1),
		Name: "test-topic",
	}

	err := repository.Add(topic)

	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}
}

func TestTopicRepositoryCannotAddDuplicateTopic(test *testing.T) {
	repository := NewTopicRepository()

	topic := topics.Topic{
		ID:   topics.TopicID(1),
		Name: "test-topic",
	}

	err := repository.Add(topic)

	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}

	err = repository.Add(topic)
	if err == nil {
		test.Fatal("Expected error for duplicate topic")
	}
}
