package handlers

import (
	"errors"
	"fmt"
	"kestrel/internal/adapters"
	"kestrel/internal/adapters/memory"
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
	"kestrel/internal/core/broker/ports"
	"testing"
)

type mockRepository struct {
	err error // что вернёт Add
}

func (s mockRepository) Add(topics.Topic) error {
	return s.err
}

func TestCreateTopicHandlerCanHandle(test *testing.T) {
	idGenerator := adapters.NewIDGenerator()
	repository := memory.NewTopicRepository()

	handler := NewCreateTopicHandler(idGenerator, repository)

	request := requests.CreateTopic{
		CorrelationID: messages.CorrelationID(1),
		TopicName:     "test-topic",
	}
	expected := responses.CreateTopicOK{CorrelationID: request.CorrelationID}

	response := handler.Handle(request)

	if response != expected {
		test.Fatalf("got %+v, expected %+v", response, expected)
	}
}

func TestCreateTopicHandlerCanNotHandleIfTopicExists(test *testing.T) {
	idGenerator := adapters.NewIDGenerator()
	repository := mockRepository{err: ports.ErrTopicAlreadyExists}

	handler := NewCreateTopicHandler(idGenerator, repository)

	request := requests.CreateTopic{
		CorrelationID: messages.CorrelationID(1),
		TopicName:     "test-topic",
	}
	expected := responses.NewFailure(
		request.CorrelationID,
		responses.CodeTopicExists,
		fmt.Sprintf("failed to add topic: %v", ports.ErrTopicAlreadyExists))

	response := handler.Handle(request)

	if response != expected {
		test.Fatalf("got %+v, expected %+v", response, expected)
	}
}

func TestCreateTopicHandlerCanNotHandleIfGetInternalError(test *testing.T) {
	idGenerator := adapters.NewIDGenerator()

	repositoryErr := errors.New("some error")
	repository := mockRepository{err: repositoryErr}

	handler := NewCreateTopicHandler(idGenerator, repository)

	request := requests.CreateTopic{
		CorrelationID: messages.CorrelationID(1),
		TopicName:     "test-topic",
	}
	expected := responses.NewFailure(
		request.CorrelationID,
		responses.CodeInternal,
		fmt.Sprintf("failed to add topic: %v", repositoryErr))

	response := handler.Handle(request)

	if response != expected {
		test.Fatalf("got %+v, expected %+v", response, expected)
	}
}
