package requests

// RequestType — тип запроса.
type RequestType uint8

const (
	TypePublish       RequestType = 0x01
	TypeSubscribe     RequestType = 0x02
	TypeUnsubscribe   RequestType = 0x03
	TypeCreateTopic   RequestType = 0x04
	TypePublishOk     RequestType = 0x81
	TypeSubscribeOk   RequestType = 0x82
	TypeUnsubscribeOk RequestType = 0x83
	TypeCreateTopicOk RequestType = 0x84
	TypeDeliver       RequestType = 0x85
	TypeError         RequestType = 0xFF
)
