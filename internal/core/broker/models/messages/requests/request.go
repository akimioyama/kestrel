package requests

import "kestrel/internal/core/broker/models/messages"

// Request — команда от клиента.
type Request interface {
	Type() messages.MessageType
}
