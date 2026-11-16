package handlers

import (
	"errors"
	"fmt"
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
	"kestrel/internal/core/broker/ports"
	"testing"
)

// mockIDGenerator всегда возвращает один и тот же ID.
type mockIDGenerator struct {
	id uint64 // что вернёт NextID
}

func (s mockIDGenerator) NextID() uint64 {
	return s.id
}

// mockRepository возвращает заданную ошибку и запоминает, что ему передали.
type mockRepository struct {
	err   error          // что вернёт Add
	added []topics.Topic // что передали в Add
}

func (s *mockRepository) Add(topic topics.Topic) error {
	s.added = append(s.added, topic)
	return s.err
}

// Проверка на этапе компиляции: если интерфейсы портов изменятся,
// тесты перестанут собираться, и мы сразу увидим, какой мок поправить.
var (
	_ ports.IDGenerator     = mockIDGenerator{}
	_ ports.TopicRepository = (*mockRepository)(nil)
)

type createTopicHandlerTestCase struct {
	testName         string
	repositoryErr    error              // что вернёт репозиторий
	expectedResponse responses.Response // что должен ответить хендлер
}

func TestCreateTopicHandlerHandle(test *testing.T) {
	request := requests.CreateTopic{
		CorrelationID: messages.CorrelationID(1),
		TopicName:     "test-topic",
	}
	internalErr := errors.New("some error")

	testCases := []createTopicHandlerTestCase{
		{
			"Topic created",
			nil,
			responses.CreateTopicOK{CorrelationID: request.CorrelationID},
		},
		{
			"Topic already exists",
			ports.ErrTopicAlreadyExists,
			responses.NewFailure(
				request.CorrelationID,
				responses.CodeTopicExists,
				fmt.Sprintf("failed to add topic: %v", ports.ErrTopicAlreadyExists)),
		},
		{
			"Repository internal error",
			internalErr,
			responses.NewFailure(
				request.CorrelationID,
				responses.CodeInternal,
				fmt.Sprintf("failed to add topic: %v", internalErr)),
		},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			// Arrange
			idGenerator := mockIDGenerator{id: 42}
			repository := &mockRepository{err: tc.repositoryErr}
			handler := NewCreateTopicHandler(idGenerator, repository)
			expectedTopic := topics.Topic{ID: topics.TopicID(42), Name: request.TopicName}

			// Act
			response := handler.Handle(request)

			// Assert
			if response != tc.expectedResponse {
				t.Fatalf("got %+v, expected %+v", response, tc.expectedResponse)
			}

			if len(repository.added) != 1 {
				t.Fatalf("Add called %d times, expected 1", len(repository.added))
			}

			if repository.added[0] != expectedTopic {
				t.Errorf("added %+v, expected %+v", repository.added[0], expectedTopic)
			}
		})
	}
}
