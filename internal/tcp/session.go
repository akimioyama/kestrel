package tcp

import (
	"io"
	"kestrel/internal/core/broker/models/messages/responses"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"sync"
)

// SessionID - идентификатор сессии.
type SessionID uint64

// Session - сессия работы с подключением.
type Session struct {
	id                  SessionID
	decoder             transport.Decoder
	encoder             transport.Encoder
	handler             ports.Handler
	readWriteCloser     io.ReadWriteCloser
	responces           chan responses.Response
	subscriberOnClosure chan<- SessionID
	closureExecutor     sync.Once
}

// NewSession создаёт новую Session.
func NewSession(id SessionID, decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler, readWriteCloser io.ReadWriteCloser) *Session {
	return &Session{
		id:                  id,
		decoder:             decoder,
		encoder:             encoder,
		handler:             handler,
		readWriteCloser:     readWriteCloser,
		responces:           make(chan responses.Response, 20), // [todo] сделать конфигурируемым параметром? что делать при переполнении очереди из-за медленной рассылки?
		subscriberOnClosure: nil,
	}
}

// ID предоставляет id сессии
func (session *Session) ID() SessionID { return session.id }

// Write пишет ответы.
func (session *Session) Write(response responses.Response) {
	select {
	case session.responces <- response: // [todo] что делать при переполнении буфера? пока сделал закрытие
		return
	default:
		go session.Close()
	}
}

// Close закрывает сессию.
func (session *Session) Close() {
	session.closureExecutor.Do(
		func() {
			session.readWriteCloser.Close()
			close(session.responces)

			go func() { session.subscriberOnClosure <- session.id }()
		})
}

// Run запускает сессию.
func (session *Session) Run(ch chan<- SessionID) {
	session.subscriberOnClosure = ch
	go session.runReader()
	go session.runWriter()

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
