package tcp

import (
	"errors"
	"kestrel/internal/core/broker/models/messages"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
	"reflect"
	"testing"
)

type mapperToDomainTestCase struct {
	testName string
	input    Frame
	error    error
	expected requests.Request
}

type mapperToTransportTestCase struct {
	testName string
	input    responses.Response
	error    error
	expected Frame
}

type fakeResponse struct {
	responses.Response
}

func TestToDomain(test *testing.T) {
	testCases := []mapperToDomainTestCase{
		// TypeCreateTopic
		// Валидный маппинг
		{
			"Can map to domain create topic type",
			Frame{
				Type: uint8(messages.TypeCreateTopic),
				Header: Header{
					CorrelationID: new(uint64(123)),
					Topic:         new("new-topic"),
				},
			},
			nil,
			requests.CreateTopic{
				CorrelationID: messages.CorrelationID(123),
				TopicName:     topics.TopicName("new-topic"),
			},
		},
		// Нет CorrelationID
		{
			"Can not map to domain create topic type if correlation id is null",
			Frame{
				Type: uint8(messages.TypeCreateTopic),
				Header: Header{
					CorrelationID: nil,
					Topic:         new("new-topic"),
				},
			},
			ErrNotFoundCorrelationID,
			nil,
		},
		// Нет Topic
		{
			"Can not map to domain create topic type if topic is null",
			Frame{
				Type: uint8(messages.TypeCreateTopic),
				Header: Header{
					CorrelationID: new(uint64(123)),
					Topic:         nil,
				},
			},
			ErrNotFoundTopic,
			nil,
		},
		// Имя топика не проходит валидацию
		{
			"Can not map to domain create topic type if topic name is not supported",
			Frame{
				Type: uint8(messages.TypeCreateTopic),
				Header: Header{
					CorrelationID: new(uint64(123)),
					Topic:         new("new topic"),
				},
			},
			topics.ErrInvalidTopicName,
			nil,
		},

		// Не поддерживаемый тип
		{
			"Can not map to domain if type is not supported",
			Frame{
				Type: 0x00,
			},
			ErrUnknownRequestType,
			nil,
		},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			// Arrange
			mapper := Mapper{}

			// Act
			result, err := mapper.ToDomain(tc.input)

			// Assert
			if !errors.Is(err, tc.error) {
				t.Errorf("ToDomain(%+v) error = %v, expected %v", tc.input, err, tc.error)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ToDomain(%+v) = %+v, expected %+v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestToTransport(test *testing.T) {
	testCases := []mapperToTransportTestCase{
		// Failure
		// Валидный маппинг
		{
			"Can map failure to transport respons",
			responses.Failure{
				CorrelationID: messages.CorrelationID(123),
				Code:          responses.CodeInternal,
				Message:       "error",
			},
			nil,
			Frame{
				Type: uint8(messages.TypeError),
				Header: Header{
					CorrelationID: new(uint64(123)),
					ErrorCode:     new(uint8(responses.CodeInternal)),
					ErrorMessage:  new("error")},
			},
		},
		// CreateTopicOK
		// Валидный маппинг
		{
			"Can map create topic ok to transport respons",
			responses.CreateTopicOK{
				CorrelationID: messages.CorrelationID(123),
			},
			nil,
			Frame{
				Type: uint8(messages.TypeCreateTopicOk),
				Header: Header{
					CorrelationID: new(uint64(123)),
				},
			},
		},
		// Не поддерживаемый тип
		{
			"Can not map to transport if type is not supported",
			fakeResponse{},
			ErrUnknownResponseType,
			Frame{},
		}}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			// Arrange
			mapper := Mapper{}

			// Act
			result, err := mapper.ToTransport(tc.input)

			// Assert
			if !errors.Is(err, tc.error) {
				t.Errorf("ToTransport(%+v) error = %v, expected %v", tc.input, err, tc.error)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ToTransport(%+v) = %+v, expected %+v", tc.input, result, tc.expected)
			}
		})
	}
}
