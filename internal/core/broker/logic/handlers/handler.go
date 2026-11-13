package handlers

import (
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
)

// Handler обрабатывает запросы.
type Handler interface {
	Handle(request requests.Request) responses.Response
}
