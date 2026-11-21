package transport

import (
	"kestrel/internal/core/broker/models/messages/responses"
)

// Encoder - конвертер доменные моделей в байт код.
type Encoder interface {
	// Encode конвертирует доменные модели ответов брокера в байт код.
	Encode(response responses.Response) ([]byte, error)
}
