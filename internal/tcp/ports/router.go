package ports

import (
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
)

type Router interface {
	Route(request requests.Request) responses.Response
}
