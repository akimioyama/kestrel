package topics

import "testing"

type topicTestCase struct {
	testName       string
	inputTopicID   TopicID
	inputTopicName TopicName
	expectValue    Topic
}

func TestNewTopic(test *testing.T) {
	testCases := []topicTestCase{
		{"Can create topic", TopicID{Value: 1}, TopicName("name"), Topic{TopicID{Value: 1}, TopicName("name")}},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			actual := NewTopic(tc.inputTopicID, tc.inputTopicName)

			// Проверяем поле InternalID по отдельности
			if actual.InternalID != tc.expectValue.InternalID {
				t.Errorf("NewTopic() InternalID = %v, expected %v", actual.InternalID, tc.expectValue.InternalID)
			}

			// Проверяем поле Name по отдельности
			if actual.Name != tc.expectValue.Name {
				t.Errorf("NewTopic() Name = %q, expected %q", actual.Name, tc.expectValue.Name)
			}
		})
	}
}
