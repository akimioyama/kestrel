package memory

import (
	"fmt"
	"sync"

	"kestrel/internal/core/broker/models/topics"
	"kestrel/internal/core/broker/ports"
)

// TopicRepository — репозиторий топиков в памяти.
type TopicRepository struct {
	topics map[topics.TopicName]topics.Topic
	mu     sync.Mutex
}

// NewTopicRepository создаёт новый репозиторий топиков в памяти.
func NewTopicRepository() *TopicRepository {
	return &TopicRepository{
		topics: make(map[topics.TopicName]topics.Topic),
	}
}

// Add добавляет топик в репозиторий. Если топик с таким именем уже есть, возвращает ошибку.
func (repository *TopicRepository) Add(topic topics.Topic) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, exists := repository.topics[topic.Name]; exists {
		return fmt.Errorf("%w: %s", ports.ErrTopicAlreadyExists, topic.Name)
	}

	repository.topics[topic.Name] = topic
	return nil

}
