package protocol

import "fmt"

// Type — тип операции протокола.
type Type uint8

const (
	// Client -> Server.
	Publish     Type = 0x01
	Subscribe   Type = 0x02
	Unsubscribe Type = 0x03
	CreateTopic Type = 0x04

	// Server -> Client: ответы на запросы.
	PublishOK     Type = 0x81
	SubscribeOK   Type = 0x82
	UnsubscribeOK Type = 0x83
	CreateTopicOK Type = 0x84

	// Server -> Client: события и ошибки.
	Deliver Type = 0x85
	Error   Type = 0xFF
)

func (t Type) String() string {
	switch t {
	case Publish:
		return "PUBLISH"
	case Subscribe:
		return "SUBSCRIBE"
	case Unsubscribe:
		return "UNSUBSCRIBE"
	case CreateTopic:
		return "CREATE_TOPIC"
	case PublishOK:
		return "PUBLISH_OK"
	case SubscribeOK:
		return "SUBSCRIBE_OK"
	case UnsubscribeOK:
		return "UNSUBSCRIBE_OK"
	case CreateTopicOK:
		return "CREATE_TOPIC_OK"
	case Deliver:
		return "DELIVER"
	case Error:
		return "ERROR"
	default:
		return fmt.Sprintf("UNKNOWN(0x%02X)", uint8(t))
	}
}
