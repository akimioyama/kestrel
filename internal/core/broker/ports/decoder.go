package ports

import (
	"io"
	"kestrel/internal/core/broker/models/messages/requests"
)

type Decoder interface {
	Decode(io.Reader) (requests.Request, error)
}
