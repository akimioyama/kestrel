package handlers

import (
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
)

// Handler - обрабатчик запросов.
type Handler interface {
	// Handle обрабатывает запросы.
	Handle(request requests.Request) responses.Response
}
