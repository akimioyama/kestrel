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
	id              SessionID
	decoder         transport.Decoder
	encoder         transport.Encoder
	handler         ports.Handler
	readWriteCloser io.ReadWriteCloser
	responces       chan responses.Response
	done            chan struct{}
	closureExecutor sync.Once
}

// NewSession создаёт новую Session.
func NewSession(id SessionID, decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler, readWriteCloser io.ReadWriteCloser) *Session {
	return &Session{
		id:              id,
		decoder:         decoder,
		encoder:         encoder,
		handler:         handler,
		readWriteCloser: readWriteCloser,
		responces:       make(chan responses.Response, 20), // [todo] сделать конфигурируемым параметром? что делать при переполнении очереди из-за медленной рассылки?
		done:            make(chan struct{}),
	}
}

// ID предоставляет id сессии
func (session *Session) ID() SessionID { return session.id }

func (session *Session) Done() <-chan struct{} { return session.done }

// Write пишет ответы.
func (session *Session) Write(response responses.Response) {
	select {
	case session.responces <- response: // [todo] что делать при переполнении буфера? пока сделал закрытие
		return
	case <-session.done:
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
			close(session.done)
		})
}

// Run запускает сессию.
func (session *Session) Run() {
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
		select {
		case <-session.done:
			return
		case session.responces <- response: // основная ветка выполнения, продолжаем цикл
		}
	}
}

func (session *Session) runWriter() {
	var response responses.Response
	for {
		select {
		case <-session.done:
			return
		case response = <-session.responces: // основная ветка выполнения, продолжаем цикл
		}

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
