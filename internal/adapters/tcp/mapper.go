package tcp

import (
	"errors"
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
)

// Mapper переводит фреймы протокола в доменные запросы и ответы обратно во фреймы.
type Mapper struct{}

// ToDomain превращает входящий фрейм в доменный запрос.
func (m *Mapper) ToDomain(frame Frame) (requests.Request, error) {
	switch messages.MessageType(frame.Type) {
	case messages.TypeCreateTopic:

		if frame.Header.CorrelationID == nil {
			return nil, errors.New("")
		}

		return requests.CreateTopic{
			CorrelationID: messages.CorrelationID(*frame.Header.CorrelationID),
			TopicName:     topics.TopicName(*frame.Header.Topic),
		}, nil
	default:
		return nil, nil
	}
}

// ToTransport превращает доменный ответ во фрейм для отправки клиенту.
func (m *Mapper) ToTransport(respons responses.Response) (Frame, error) {
	switch r := respons.(type) {
	case responses.CreateTopicOK:
		correlationID := uint64(r.CorrelationID)
		return Frame{
			Type:   uint8(messages.TypeCreateTopicOk),
			Header: Header{CorrelationID: &correlationID},
		}, nil
	default:
		return Frame{}, nil
	}
}
