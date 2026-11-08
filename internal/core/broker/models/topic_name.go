package models

import (
	"errors"
	"regexp"
)

// TopicName — валидное имя топика.
type TopicName string

var (
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

func (t TopicName) String() string {
	return string(t)
}
