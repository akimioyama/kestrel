package responses

import (
	"testing"

	"kestrel/internal/core/broker/models/commands"
)

type createTopicOKTestCase struct {
	testName           string
	inputCorrelationID commands.CorrelationID
}

func TestNewCreateTopicOK(test *testing.T) {
	testCases := []createTopicOKTestCase{
		{
			"Create topic OK response",
			commands.CorrelationID(123),
		},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			actual := NewCreateTopicOK(tc.inputCorrelationID)

			// Проверяем поле CorrelationID по отдельности
			if actual.CorrelationID != tc.inputCorrelationID {
				t.Errorf("NewCreateTopicOK() CorrelationID = %v, expected %v", actual.CorrelationID, tc.inputCorrelationID)
			}
		})
	}
}
