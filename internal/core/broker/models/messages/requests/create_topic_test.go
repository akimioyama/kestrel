package requests

import (
	"kestrel/internal/core/broker/models/messages"
	"testing"
)

func TestCreateTopicType(test *testing.T) {
	// Arrange
	request := CreateTopic{}

	// Act
	actual := request.Type()

	// Assert
	if actual != messages.TypeCreateTopic {
		test.Errorf("CreateTopic.Type() = %v, want %v", actual, messages.TypeCreateTopic)
	}
}
