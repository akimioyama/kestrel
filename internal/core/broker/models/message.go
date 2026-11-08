package models

import "time"

// Message — внутренняя модель сообщения, циркулирующего в брокере.
type Message struct {
	ID        uint64    // Глобальный ID, присваивается сервером
	TopicName TopicName // Имя топика, в котором опубликовано сообщение
	Headers   []byte    // Хедеры пользователя для потребителя
	Key       []byte    // Ключ
	Payload   []byte    // Тело сообщения (может быть пустым)
	Timestamp time.Time // Время приема сообщения сервером
}
