package requests

import (
	"testing"

	"kestrel/internal/core/broker/models/commands"
	"kestrel/internal/core/broker/models/topics"
)

type createTopicTestCase struct {
	testName           string
	inputCorrelationID commands.CorrelationID
	inputTopicName     topics.TopicName
	expectType         RequestType
}

func TestNewCreateTopic(test *testing.T) {
	testCases := []createTopicTestCase{
		{
			"Create topic request",
			commands.CorrelationID(123),
			topics.TopicName("my-topic"),
			TypeCreateTopic,
		},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			// Тестируем конструктор NewCreateTopic
			actual := NewCreateTopic(tc.inputCorrelationID, tc.inputTopicName)

			// Проверяем корректность заполнения полей структуры по аналогии
			if actual.CorrelationID != tc.inputCorrelationID {
				t.Errorf("NewCreateTopic() CorrelationID = %v, expected %v", actual.CorrelationID, tc.inputCorrelationID)
			}

			if actual.TopicName != tc.inputTopicName {
				t.Errorf("NewCreateTopic() TopicName = %v, expected %v", actual.TopicName, tc.inputTopicName)
			}

			// Тестируем метод Type()
			if actual.Type() != tc.expectType {
				t.Errorf("actual.Type() = %v, expected %v", actual.Type(), tc.expectType)
			}
		})
	}
}
