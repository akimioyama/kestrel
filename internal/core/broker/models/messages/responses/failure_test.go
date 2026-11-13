package responses

import (
	"kestrel/internal/core/broker/models/messages"
	"testing"
)

type failureTestCase struct {
	testName    string
	input       string
	expectPanic bool
}

func TestNewFailure(test *testing.T) {
	testCases := []failureTestCase{
		// Валидные сообщения
		{"Just text", "some message with error", false},

		// Невалидные сообщения
		{"Empty string", "", true},
		{"Only whitespace", "     ", true},
		{"Tabs and newlines", "\t\n", true},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			id := messages.CorrelationID(1)
			code := CodeInternal

			// Паника прерывает подтест, поэтому ловим её в defer:
			// он выполнится и при панике, и при обычном завершении.
			defer func() {
				r := recover()
				panicked := r != nil

				if panicked != tc.expectPanic {
					t.Errorf("NewFailure(%q) panicked = %v (%v), expectPanic %v", tc.input, panicked, r, tc.expectPanic)
				}
			}()

			actual := NewFailure(id, code, tc.input)

			// Сюда доходим, только если паники не было — проверяем результат
			if actual.Message != tc.input {
				t.Errorf("NewFailure(%q) Message = %q, expected %q", tc.input, actual.Message, tc.input)
			}

			if actual.CorrelationID != id {
				t.Errorf("NewFailure(%q) CorrelationID = %v, expected %v", tc.input, actual.CorrelationID, id)
			}

			if actual.Code != code {
				t.Errorf("NewFailure(%q) Code = %v, expected %v", tc.input, actual.Code, code)
			}
		})
	}
}
