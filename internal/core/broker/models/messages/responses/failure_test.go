package responses

import (
	"kestrel/internal/core/broker/models/messages"
	"testing"
)

type failureTestCase struct {
	testName    string
	input       string
	expectError bool
}

func TestNewFailure(test *testing.T) {
	testCases := []failureTestCase{
		// Валидные сообщения
		{"Just text", "some message with error", false},

		// Невалидные сообщения
		{"Empty string", "", true},
		{"Only whitespace", "     ", true},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			id := messages.CorrelationID(1)
			code := CodeInternal

			actual, err := NewFailure(id, code, tc.input)

			// Проверяем, совпадает ли факт наличия ошибки с ожиданием
			hasError := (err != nil)
			if hasError != tc.expectError {
				t.Fatalf("NewFailure(%q) error = %v, expectError %v", tc.input, err, tc.expectError)
				return
			}

			// Если мы ждали ошибку и она действительно вернулась, выходим из подтеста
			if tc.expectError {
				return
			}

			// Если ошибки нет, проверяем результат
			if actual.Message != tc.input {
				t.Errorf("NewFailure(%q) = %v, expected %v", tc.input, actual, tc.input)
				return
			}

			if actual.CorrelationID != id {
				t.Errorf("NewFailure(%q) CorrelationID = %v, expected %v", tc.input, actual.CorrelationID, id)
				return
			}

			if actual.Code != code {
				t.Errorf("NewFailure(%q) Code = %v, expected %v", tc.input, actual.Code, code)
				return
			}
		})
	}
}
