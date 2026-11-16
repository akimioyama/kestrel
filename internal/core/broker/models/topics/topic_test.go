package topics

import "testing"

func TestNewTopic(test *testing.T) {
	// Arrange
	id := TopicID(42)
	name := TopicName("orders")
	expected := Topic{ID: id, Name: name}

	// Act
	actual := NewTopic(id, name)

	// Assert
	// Структуры из сравнимых полей можно сравнивать через != целиком
	if actual != expected {
		test.Errorf("NewTopic(%v, %q) = %+v, expected %+v", id, name, actual, expected)
	}
}
