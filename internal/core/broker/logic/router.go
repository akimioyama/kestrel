package logic

import (
	"fmt"
	"kestrel/internal/core/broker/logic/handlers"
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
)

// Router используется для перенаправления запроса обработчикам
type Router struct {
	handlers map[messages.MessageType]handlers.Handler
}

// Route перенаправляет запросы обработчикам
func (router *Router) Route(request requests.Request) responses.Response {
	handler, ok := router.handlers[request.Type()]

	if !ok {
		panic(
			fmt.Sprintf(
				"Unexpected message type %v",
				request.Type()))
	}

	return handler.Handle(request)
}
