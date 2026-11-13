package ports

// IDGenerator выдаёт уникальные идентификаторы.
type IDGenerator interface {
	// NextID возвращает новый ID, который ещё не выдавался.
	NextID() uint64
}
