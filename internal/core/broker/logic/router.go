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

func NewRouter(handlers map[messages.MessageType]handlers.Handler) Router {
	return Router{handlers: handlers}
}

// Route перенаправляет запросы обработчикам
func (router *Router) Route(request requests.Request) responses.Response { // мб переименовать в Handle? а то в сессии и везде выше логично было бы, если был обработчик, а не роутер. Назвал так из-за тогго, что в задаче так было.
	handler, ok := router.handlers[request.Type()]

	if !ok {
		panic(
			fmt.Sprintf(
				"Unexpected message type %v",
				request.Type()))
	}

	return handler.Handle(request)
}
