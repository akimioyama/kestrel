package messages

// MessageType — тип запроса.
type MessageType uint8

const (
	TypePublish       MessageType = 0x01
	TypeSubscribe     MessageType = 0x02
	TypeUnsubscribe   MessageType = 0x03
	TypeCreateTopic   MessageType = 0x04
	TypePublishOk     MessageType = 0x81
	TypeSubscribeOk   MessageType = 0x82
	TypeUnsubscribeOk MessageType = 0x83
	TypeCreateTopicOk MessageType = 0x84
	TypeDeliver       MessageType = 0x85
	TypeError         MessageType = 0xFF
)
