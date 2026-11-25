package tcp

import (
	"errors"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
	"sync"
)

// ErrManagerClosed о том, что менеджер сессий уже закрыт.
var ErrManagerClosed = errors.New("Session manager closed")

// ErrSessionNotFound о том, что сессия не найдена.
var ErrSessionNotFound = errors.New("Session not found")

// SessionManager - управляющий сессиями.
type SessionManager struct {
	sessions             map[SessionID]*Session
	topicSubscribers     map[topics.TopicID]map[SessionID]*Session
	sessionSubscriptions map[SessionID]map[topics.TopicID]bool
	isClosed             bool
	mu                   sync.RWMutex
}

// NewSessionManager создаёт новый SessionManager.
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions:             make(map[SessionID]*Session),
		topicSubscribers:     make(map[topics.TopicID]map[SessionID]*Session),
		sessionSubscriptions: make(map[SessionID]map[topics.TopicID]bool),
	}
}

// AddSession добавляет сессию.
func (manager *SessionManager) AddSession(session *Session) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	manager.sessions[session.ID()] = session
	go manager.runCleaner(session)

	session.Run()
}

// SubscribeSessionOnTopic добавляет подписку сессии на топик.
func (manager *SessionManager) SubscribeSessionOnTopic(sessionID SessionID, topicID topics.TopicID) error { // [todo] сделать снос топиков ещё отдельным методом?
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if manager.isClosed {
		return ErrManagerClosed
	}

	session, ok := manager.sessions[sessionID]

	if !ok {
		return ErrSessionNotFound // case: пока отрабатывал хендлер сессия разорвалась
	}

	sessionSubscriptions, ok := manager.sessionSubscriptions[sessionID]
	if !ok {
		sessionSubscriptions = make(map[topics.TopicID]bool)
		manager.sessionSubscriptions[sessionID] = sessionSubscriptions
	}
	sessionSubscriptions[topicID] = true

	topicSubscribers, ok := manager.topicSubscribers[topicID]
	if !ok {
		topicSubscribers = make(map[SessionID]*Session)
		manager.topicSubscribers[topicID] = topicSubscribers
	}
	topicSubscribers[sessionID] = session

	return nil
}

// Delivery доставляет сообщения топика всем подписанным сессиям.
func (manager *SessionManager) Delivery(topicSource topics.TopicID, response responses.Response) { // remark: порядок сообщений в конкретном топике должен гарантироваться вызывающим. Например, с помощью map[TopicID]chan Responses
	manager.mu.RLock()
	defer manager.mu.RUnlock()

	for _, session := range manager.topicSubscribers[topicSource] {
		session.Write(response)
	}
}

// CloseAll закрывает все текущие и будущие сессии.
func (manager *SessionManager) CloseAll() {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	manager.isClosed = true

	for _, session := range manager.sessions {
		session.Close()
	}
}

func (manager *SessionManager) runCleaner(session *Session) {
	for {
		<-session.Done()
		sessionID := session.ID()

		manager.mu.Lock()

		delete(manager.sessions, sessionID)
		for topic := range manager.sessionSubscriptions[sessionID] {
			delete(manager.topicSubscribers[topic], sessionID)
		}
		delete(manager.sessionSubscriptions, sessionID)

		manager.mu.Unlock()
	}
}
