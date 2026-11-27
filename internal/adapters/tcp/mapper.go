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

// Ошибки, которые возвращает Mapper при маппинге.
var (
	ErrNotFoundCorrelationID = errors.New("correlationID not found")
	ErrNotFoundTopic         = errors.New("topic not found")
	ErrUnknownRequestType    = errors.New("unknown request type")
	ErrUnknownResponseType   = errors.New("unknown response type")
)

// ToDomain превращает входящий фрейм в доменный запрос.
func (m *Mapper) ToDomain(frame Frame) (requests.Request, error) {
	switch messages.MessageType(frame.Type) {
	case messages.TypeCreateTopic:
		if frame.Header.CorrelationID == nil {
			return nil, ErrNotFoundCorrelationID
		}
		if frame.Header.Topic == nil {
			return nil, ErrNotFoundTopic
		}

		name, err := topics.NewTopicName(*frame.Header.Topic)
		if err != nil {
			return nil, err
		}

		return requests.CreateTopic{
			CorrelationID: messages.CorrelationID(*frame.Header.CorrelationID),
			TopicName:     name,
		}, nil
	// новые case запросы далее по ходу добавления
	default:
		return nil, ErrUnknownRequestType
	}
}

// ToTransport превращает доменный ответ во фрейм для отправки клиенту.
func (m *Mapper) ToTransport(respons responses.Response) (Frame, error) {
	switch r := respons.(type) {
	case responses.Failure:
		correlationID := uint64(r.CorrelationID)
		return Frame{
			Type: uint8(messages.TypeError),
			Header: Header{
				CorrelationID: &correlationID,
				ErrorCode:     (*uint8)(&r.Code),
				ErrorMessage:  &r.Message},
		}, nil
	case responses.CreateTopicOK:
		correlationID := uint64(r.CorrelationID)
		return Frame{
			Type:   uint8(messages.TypeCreateTopicOk),
			Header: Header{CorrelationID: &correlationID},
		}, nil
	// новые case ответы далее по ходу добавления
	default:
		return Frame{}, ErrUnknownResponseType
	}
}
