package memory

import (
	"errors"
	"kestrel/internal/core/broker/models/topics"
	"kestrel/internal/core/broker/ports"
	"maps"
	"testing"
)

type topicRepositoryAddTestCase struct {
	testName       string
	existing       []topics.Topic                    // что уже лежит в репозитории
	topic          topics.Topic                      // что добавляем
	expectedErr    error                             // nil — ошибки быть не должно
	expectedTopics map[topics.TopicName]topics.Topic // что должно оказаться в map
}

func TestTopicRepositoryAdd(test *testing.T) {
	orders := topics.Topic{ID: 1, Name: "orders"}
	payments := topics.Topic{ID: 2, Name: "payments"}
	// То же имя, что у orders, но другой ID — так увидим, если Add перезапишет топик
	ordersDuplicate := topics.Topic{ID: 3, Name: orders.Name}

	testCases := []topicRepositoryAddTestCase{
		{
			testName:       "Add to empty repository",
			existing:       nil,
			topic:          orders,
			expectedErr:    nil,
			expectedTopics: map[topics.TopicName]topics.Topic{orders.Name: orders},
		},
		{
			testName:       "Add topic with another name",
			existing:       []topics.Topic{orders},
			topic:          payments,
			expectedErr:    nil,
			expectedTopics: map[topics.TopicName]topics.Topic{orders.Name: orders, payments.Name: payments},
		},
		{
			testName:       "Add duplicate name",
			existing:       []topics.Topic{orders},
			topic:          ordersDuplicate,
			expectedErr:    ports.ErrTopicAlreadyExists,
			expectedTopics: map[topics.TopicName]topics.Topic{orders.Name: orders},
		},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			// Arrange
			repository := NewTopicRepository()

			for _, topic := range tc.existing {
				if err := repository.Add(topic); err != nil {
					t.Fatalf("Expected no error on preparing Add(%+v), got %v", topic, err)
				}
			}

			// Act
			err := repository.Add(tc.topic)

			// Assert
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Add(%+v) error = %v, expected %v", tc.topic, err, tc.expectedErr)
			}

			if !maps.Equal(repository.topics, tc.expectedTopics) {
				t.Errorf("topics = %v, expected %v", repository.topics, tc.expectedTopics)
			}
		})
	}
}
