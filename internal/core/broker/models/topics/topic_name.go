package topics

import (
	"errors"
	"regexp"
)

// TopicName — валидированное имя топика.
type TopicName string

var (
	// ErrInvalidTopicName возвращается, если имя топика не соответствует правилам именования.
	ErrInvalidTopicName = errors.New("invalid topic name")

	// Правило: [a-z0-9._-], 1-255 символов, не начинается с . или -
	topicNameRegex = regexp.MustCompile(`^[a-z0-9_][a-z0-9._-]{0,254}$`)
)

// NewTopicName создает новый TopicName, проверяя его на валидность.
func NewTopicName(name string) (TopicName, error) {
	if !topicNameRegex.MatchString(name) {
		return "", ErrInvalidTopicName
	}
	return TopicName(name), nil
}

// String возвращает строковое представление имени топика.
func (t TopicName) String() string {
	return string(t)
}
