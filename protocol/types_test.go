package protocol

import "testing"

type TestCase struct {
	testName      string
	operationType Type
	expected      string
}

func TestTypeString(t *testing.T) {
	tests := []TestCase{
		{"Publish", Publish, "PUBLISH"},
		{"Subscribe", Subscribe, "SUBSCRIBE"},
		{"Unsubscribe", Unsubscribe, "UNSUBSCRIBE"},
		{"CreateTopic", CreateTopic, "CREATE_TOPIC"},
		{"PublishOK", PublishOK, "PUBLISH_OK"},
		{"SubscribeOK", SubscribeOK, "SUBSCRIBE_OK"},
		{"UnsubscribeOK", UnsubscribeOK, "UNSUBSCRIBE_OK"},
		{"CreateTopicOK", CreateTopicOK, "CREATE_TOPIC_OK"},
		{"Deliver", Deliver, "DELIVER"},
		{"Error", Error, "ERROR"},
		{"Unknown", Type(0x99), "UNKNOWN(0x99)"},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			actual := tc.operationType.String()
			if actual != tc.expected {
				t.Errorf("Type.String() = %v, expected %v", actual, tc.expected)
			}
		})
	}
}
