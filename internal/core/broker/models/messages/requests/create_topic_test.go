package requests

import (
	"kestrel/internal/core/broker/models/messages"
	"testing"
)

func TestCreateTopicType(test *testing.T) {
	request := CreateTopic{}

	actual := request.Type()

	if actual != messages.TypeCreateTopic {
		test.Errorf("CreateTopic.Type() = %v, want %v", actual, messages.TypeCreateTopic)
	}
}
