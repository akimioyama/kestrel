package topics

import "testing"

type topicIdTestCase struct {
	testName    string
	input       uint64
	expectValue uint64
}

func TestNewTopicId(test *testing.T) {
	testCases := []topicIdTestCase{
		{"Create topic id", 12, 12},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			actual := NewTopicID(tc.input)

			// Если ошибки нет, проверяем результат
			if actual.Value != tc.input {
				t.Errorf("NewTopicID(%d) = %v, expected %v", tc.input, actual, tc.input)
			}
		})
	}
}
