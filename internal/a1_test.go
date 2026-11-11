package internal

import (
	"testing"
)

type aTestCase struct {
	testName string
	input    int32
}

func TestNewCreateTopic(test *testing.T) {
	testCases := []aTestCase{
		{
			"Chech \"a\" convertation",
			2,
		},
	}

	for _, tc := range testCases {
		test.Run(tc.testName, func(t *testing.T) {
			// Тестируем конструктор NewCreateTopic
			actual := A1(tc.input)

			// Проверяем корректность заполнения полей структуры по аналогии
			if int32(actual) != tc.input {
				t.Errorf("actual = %v, expected %v", int32(actual), tc.input)
			}
		})
	}
}
