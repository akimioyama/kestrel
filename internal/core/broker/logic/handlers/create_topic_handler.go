package handlers

import (
	"errors"
	"fmt"

	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
	"kestrel/internal/core/broker/ports"
)

// CreateTopicHandler обрабатывает запросы на создание топиков.
type CreateTopicHandler struct {
	idGenerator ports.IDGenerator
	repository  ports.TopicRepository
}

// NewCreateTopicHandler создаёт новый обработчик запросов на создание топиков.
func NewCreateTopicHandler(idGenerator ports.IDGenerator, repository ports.TopicRepository) *CreateTopicHandler {
	return &CreateTopicHandler{
		idGenerator: idGenerator,
		repository:  repository,
	}
}

// Handle создаёт топик и сохраняет его в репозиторий.
func (handler *CreateTopicHandler) Handle(request requests.Request) responses.Response {

	createTopicRequest, ok := request.(requests.CreateTopic)

	if !ok {
		panic(
			fmt.Sprintf(
				"CreateTopicHandler: unexpected request type %T",
				request))
	}

	newTopic := topics.NewTopic(
		topics.TopicID(handler.idGenerator.NextID()),
		createTopicRequest.TopicName,
	)
	err := handler.repository.Add(newTopic)

	if err != nil {

		// Ожидаемый отказ — топик уже есть; всё остальное — сбой сервера.
		code := responses.CodeInternal

		if errors.Is(err, ports.ErrTopicAlreadyExists) {
			code = responses.CodeTopicExists
		}

		return responses.NewFailure(
			createTopicRequest.CorrelationID,
			code,
			fmt.Sprintf("failed to add topic: %v", err))
	}

	return responses.CreateTopicOK{
		CorrelationID: createTopicRequest.CorrelationID,
	}
}
