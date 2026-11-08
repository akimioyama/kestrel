package models

import (
	"errors"
	"strings"
	"testing"
)

type TopicTestCase struct {
	testName    string
	input       string
	expectError bool
}

func TestNewTopicName(t *testing.T) {
	tests := []TopicTestCase{
		// Валидные имена
		{"Simple lower", "orders", false},
		{"With underscore", "my_topic", false},
		{"With dot and dash", "my-topic.v1", false},
		{"With digits", "topic123", false},
		{"Max length (255)", strings.Repeat("a", 255), false},

		// Невалидные имена
		{"Empty string", "", true},
		{"Uppercase", "Orders", true},
		{"Starts with dot", ".hidden", true},
		{"Starts with dash", "-topic", true},
		{"Spaces", "my topic", true},
		{"Special chars", "topic!", true},
		{"Too long (256)", strings.Repeat("a", 256), true},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual, err := NewTopicName(tc.input)

			// Проверяем, совпадает ли факт наличия ошибки с ожиданием
			hasError := (err != nil)
			if hasError != tc.expectError {
				t.Fatalf("NewTopicName(%q) error = %v, expectError %v", tc.input, err, tc.expectError)
			}

			// Если ошибка ожидалась, проверяем, что она та самая
			if tc.expectError {
				if !errors.Is(err, ErrInvalidTopicName) {
					t.Errorf("NewTopicName(%q) error = %v, expected %v", tc.input, err, ErrInvalidTopicName)
				}
				return // дальше проверять нечего
			}

			// Если ошибки нет, проверяем результат
			if string(actual) != tc.input {
				t.Errorf("NewTopicName(%q) = %v, expected %v", tc.input, actual, tc.input)
			}
		})
	}
}

func TestTopicNameString(t *testing.T) {
	topicName := TopicName("orders")
	expected := "orders"

	if actual := topicName.String(); actual != expected {
		t.Errorf("TopicName.String() = %v, expected %v", actual, expected)
	}
}
