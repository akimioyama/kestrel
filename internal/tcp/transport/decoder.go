package transport

import (
	"io"
	"kestrel/internal/core/broker/models/messages/requests"
)

// Decoder - конвертер байт код в доменные модели.
type Decoder interface {
	// Decode конвертирует байт код в доменные модели запросов брокера.
	Decode(reader io.Reader) (requests.Request, error)
}
