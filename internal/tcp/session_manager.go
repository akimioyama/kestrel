package tcp

import (
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/core/broker/models/topics"
	"sync"
)

// SessionManager - управляющий сессиями.
type SessionManager struct {
	sessions             map[SessionID]*Session
	topicSubscribers     map[topics.TopicID]map[SessionID]*Session
	sessionSubscriptions map[SessionID]map[topics.TopicID]bool
	sessionCloses        chan SessionID
	isClosed             bool
	mu                   sync.RWMutex
}

// NewSessionManager создаёт новый SessionManager.
func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessions:             make(map[SessionID]*Session),
		topicSubscribers:     make(map[topics.TopicID]map[SessionID]*Session),
		sessionSubscriptions: make(map[SessionID]map[topics.TopicID]bool),
		sessionCloses:        make(chan SessionID)}
}

// Init инициализирует объект.
func (manager *SessionManager) Init() {
	go manager.runCleaner()
}

// AddSession добавляет сессию.
func (manager *SessionManager) AddSession(session *Session) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if manager.isClosed {
		session.Close() // case: сессия могла запускаться во время закрытия менеджера.
		return
	}

	manager.sessions[session.ID()] = session
	session.SubscribeOnClosure(manager.sessionCloses)
	//[todo] можно:
	// 1) заменить на Run, но нет гарантий, что Run не будет вызван дважды. с SubscribeOnClosure та же проблема
	// 2.а) сессия может тут создаваться. Но тогда слишком много всего в текущей структуре будет, а она уже не простая
	// 2.б) фабрику? но это оверинжинеринг и не нужно
}

// SubscribeSessionOnTopic добавляет подписку сессии на топик.
func (manager *SessionManager) SubscribeSessionOnTopic(sessionID SessionID, topicID topics.TopicID) { // [todo] сделать снос топиков ещё отдельным методом?
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if manager.isClosed {
		panic("Session meneger closed.")
	}

	session, ok := manager.sessions[sessionID]

	if !ok {
		panic("Session not found.") // case: пока отрабатывал хендлер сессия разорвалась
	}

	manager.sessionSubscriptions[sessionID][topicID] = true
	manager.topicSubscribers[topicID][sessionID] = session
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

	for _, session := range manager.sessions {
		session.Close()
	}

	if manager.isClosed {
		manager.isClosed = true
	}
}

func (manager *SessionManager) runCleaner() {
	for {
		session := <-manager.sessionCloses
		manager.mu.Lock()

		delete(manager.sessions, session)
		for topic := range manager.sessionSubscriptions[session] {
			delete(manager.topicSubscribers[topic], session)
		}
		delete(manager.sessionSubscriptions, session)

		manager.mu.Unlock()
	}
}
