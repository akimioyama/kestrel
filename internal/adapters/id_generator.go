package adapters

import (
	"sync"
)

// IDGenerator — генератор уникальных ID.
type IDGenerator struct {
	mu    sync.Mutex
	count uint64
}

// NewIDGenerator создаёт новый генератор ID.
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// NextID возвращает следующий уникальный ID.
func (generator *IDGenerator) NextID() uint64 {
	generator.mu.Lock()
	defer generator.mu.Unlock()

	generator.count++
	return generator.count
}
