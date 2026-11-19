package handlers

import (
	"fmt"
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
)

// Router используется для перенаправления запроса обработчикам
type Router struct {
	handlers map[messages.MessageType]Handler
}

func NewRouter(handlers map[messages.MessageType]Handler) Router {
	return Router{handlers: handlers}
}

// Handle перенаправляет запрос нужному обработчику
func (router *Router) Handle(request requests.Request) responses.Response {
	handler, ok := router.handlers[request.Type()]

	if !ok {
		panic(fmt.Sprintf("Unexpected message type %v", request.Type()))
	}

	return handler.Handle(request)
}
