package memory

import (
	"errors"
	"kestrel/internal/core/broker/models/topics"
	"kestrel/internal/core/broker/ports"
	"testing"
)

func TestTopicRepositoryCanAddTopic(test *testing.T) {
	// Arrange
	repository := NewTopicRepository()
	topic := topics.Topic{
		ID:   topics.TopicID(1),
		Name: "test-topic",
	}

	// Act
	err := repository.Add(topic)

	// Assert
	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}
}

func TestTopicRepositoryCannotAddDuplicateTopic(test *testing.T) {
	// Arrange
	repository := NewTopicRepository()
	topic := topics.Topic{
		ID:   topics.TopicID(1),
		Name: "test-topic",
	}

	// Первый Add — предусловие, а не проверяемое действие
	if err := repository.Add(topic); err != nil {
		test.Fatalf("Expected no error on first Add, got %v", err)
	}

	// Act
	err := repository.Add(topic)

	// Assert
	// errors.Is(nil, ...) возвращает false, так что nil тоже поймаем
	if !errors.Is(err, ports.ErrTopicAlreadyExists) {
		test.Fatalf("Expected %v, got %v", ports.ErrTopicAlreadyExists, err)
	}
}
