package ports

import (
	"errors"

	"kestrel/internal/core/broker/models/topics"
)

// ErrTopicAlreadyExists ошибка о том, что топик с таким именем уже есть.
var ErrTopicAlreadyExists = errors.New("topic already exists")

// TopicRepository хранит топики.
type TopicRepository interface {
	// Add добавляет топик в репозиторий.
	// Если топик с таким именем уже есть, возвращает ErrTopicAlreadyExists.
	Add(topic topics.Topic) error
}
