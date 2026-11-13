package topics

import "testing"

func TestNewTopic(test *testing.T) {
	id := TopicID(42)
	name := TopicName("orders")

	actual := NewTopic(id, name)
	expected := Topic{ID: id, Name: name}

	// Структуры из сравнимых полей можно сравнивать через != целиком
	if actual != expected {
		test.Errorf("NewTopic(%v, %q) = %+v, expected %+v", id, name, actual, expected)
	}
}
