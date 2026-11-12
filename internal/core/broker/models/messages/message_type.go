package messages

// MessageType — тип запроса.
type MessageType uint8

const (
	// TypePublish — запрос на публикацию сообщения в топик.
	TypePublish MessageType = 0x01

	// TypeSubscribe — запрос на подписку на топик.
	TypeSubscribe MessageType = 0x02

	// TypeUnsubscribe — запрос на отписку от топика.
	TypeUnsubscribe MessageType = 0x03

	// TypeCreateTopic — запрос на создание нового топика.
	TypeCreateTopic MessageType = 0x04

	// TypePublishOk — подтверждение успешной публикации.
	TypePublishOk MessageType = 0x81

	// TypeSubscribeOk — подтверждение успешной подписки.
	TypeSubscribeOk MessageType = 0x82

	// TypeUnsubscribeOk — подтверждение успешной отписки.
	TypeUnsubscribeOk MessageType = 0x83

	// TypeCreateTopicOk — подтверждение успешного создания топика.
	TypeCreateTopicOk MessageType = 0x84

	// TypeDeliver — доставка сообщения подписчику.
	TypeDeliver MessageType = 0x85

	// TypeError — сообщение об ошибке.
	TypeError MessageType = 0xFF
)
