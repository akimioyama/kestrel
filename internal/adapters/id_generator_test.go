package adapters

import (
	"sync"
	"testing"
)

func TestIDGeneratorSequentialIDs(test *testing.T) {
	generator := NewIDGenerator()
	for i := uint64(1); i <= 10; i++ {
		id := generator.NextID()
		if id != i {
			test.Errorf("Expected ID %d, got %d", i, id)
		}
	}
}

func TestIDGeneratorConcurrentAccess(test *testing.T) {
	generator := NewIDGenerator()
	numGoroutines := 100

	ids := make([]uint64, numGoroutines)
	var wg sync.WaitGroup

	for i := range numGoroutines {
		wg.Go(func() { ids[i] = generator.NextID() }) // каждая горутина пишет в свою ячейку
	}
	wg.Wait()

	seen := make(map[uint64]bool)
	for _, id := range ids {
		if seen[id] {
			test.Fatalf("duplicate ID %d", id)
		}
		seen[id] = true
	}
}
