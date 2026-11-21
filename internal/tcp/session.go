package tcp

import (
	"io"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"sync"
)

type SessionId uint64

// Session - сессия работы с подключением.
type Session struct {
	id                   SessionId
	decoder              transport.Decoder
	encoder              transport.Encoder
	handler              ports.Handler
	readWriteCloser      io.ReadWriteCloser
	responces            chan responses.Response
	subscribersOnClosure map[string]chan *Session
	isClosed             bool
	closeMutex           sync.Mutex
}

// NewSession создаёт новую Session.
func NewSession(id SessionId, decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler, readWriteCloser io.ReadWriteCloser) *Session {
	return &Session{
		id:                   id,
		decoder:              decoder,
		encoder:              encoder,
		handler:              handler,
		readWriteCloser:      readWriteCloser,
		responces:            make(chan responses.Response), // [todo] один медленный tcp будет блочить другие? при подписке несколько топиков ситуация ещё больше ухудшится
		subscribersOnClosure: make(map[string]chan *Session),
	}
}

// Id предоставляет id сессии
func (session *Session) Id() SessionId { return session.id }

// Write пишет ответы.
func (session *Session) Write(response responses.Response) { // [todo] нужен ли тут какой-то возврат ошибки? а то тут и паника может быть после закрытия канала. Ловить и конвертить в ошику?
	session.responces <- response
}

// SubscribeOnClosure сохраняет подписчика на закрытие сессии.
func (session *Session) SubscribeOnClosure(subscriber string, ch chan *Session) { // [todo] может быть тут вместо строковых имён можно добавлять айдишники топиков, но что тогда писать серверу?
	if !session.isClosed {
		session.closeMutex.Lock()
		defer session.closeMutex.Unlock()

		if !session.isClosed {
			session.subscribersOnClosure[subscriber] = ch
			return
		}
	}

	go session.notify(ch)
}

// SubscribeOnClosure удаляет подписчика на закрытие сессии.
func (session *Session) UnsubscribeOnClosure(subscriber string) { // будет использоваться при удалении топиков, например.
	if !session.isClosed {
		session.closeMutex.Lock()
		defer session.closeMutex.Unlock()

		if !session.isClosed {
			if _, ok := session.subscribersOnClosure[subscriber]; ok {
				delete(session.subscribersOnClosure, subscriber)
				return
			}
		}
	}
}

// Close закрывает сессию.
func (session *Session) Close() {
	if !session.isClosed {
		session.closeMutex.Lock()
		defer session.closeMutex.Unlock()

		if !session.isClosed {
			session.readWriteCloser.Close()
			session.isClosed = true
			close(session.responces)

			for _, ch := range session.subscribersOnClosure {
				go session.notify(ch) // [todo] go для защиты от закрытых каналов и взаимоблокировок в таких местах
			}
		}
	}
}

// Run запускает сессию.
func (session *Session) Run() {
	go session.runReader()
	go session.runWriter()
}

func (session *Session) notify(ch chan *Session) {
	ch <- session
}

func (session *Session) runReader() {
	for {
		reqest, err := session.decoder.Decode(session.readWriteCloser)

		if err != nil {
			session.Close()
			return
		}

		response := session.handler.Handle(reqest)

		session.responces <- response
	}
}

func (session *Session) runWriter() {
	for {
		response := <-session.responces

		encodedResponce, err := session.encoder.Encode(response)

		if err != nil {
			session.Close()
			return
		}

		_, err = session.readWriteCloser.Write(encodedResponce)

		if err != nil {
			session.Close()
			return
		}
	}
}
