package adapters

import (
	"sync"
	"testing"
)

// assertUniqueIDs падает, если в ids есть повторы.
// test.Helper() говорит Go, что при падении показывать строку
// в вызывающем тесте, а не здесь.
func assertUniqueIDs(test *testing.T, ids []uint64) {
	test.Helper()

	seen := make(map[uint64]bool)
	for _, id := range ids {
		if seen[id] {
			test.Fatalf("duplicate ID %d", id)
		}
		seen[id] = true
	}
}

func TestIDGeneratorUniqueIDs(test *testing.T) {
	// Arrange
	generator := NewIDGenerator()
	ids := make([]uint64, 100)

	// Act
	for i := range ids {
		ids[i] = generator.NextID()
	}

	// Assert
	assertUniqueIDs(test, ids)
}

func TestIDGeneratorConcurrentAccess(test *testing.T) {
	// Arrange
	generator := NewIDGenerator()
	numGoroutines := 100
	ids := make([]uint64, numGoroutines)
	var wg sync.WaitGroup

	// Act
	for i := range numGoroutines {
		wg.Go(func() { ids[i] = generator.NextID() }) // каждая горутина пишет в свою ячейку
	}
	wg.Wait()

	// Assert
	assertUniqueIDs(test, ids)
}
