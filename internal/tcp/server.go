package tcp

import (
	"io"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"net"
	"sync"
)

// Server - обработчик подключений.
type Server struct {
	listener           net.Listener
	decoder            transport.Decoder
	encoder            transport.Encoder
	handler            ports.Handler
	sessionIDGenerator ports.IDGenerator
	sessionsClosures   chan *Session
	sessions           map[SessionID]*Session
	muForSessions      sync.Mutex
	isStopped          bool
}

// NewServer создаёт новый NewServer.
func NewServer(listener net.Listener, decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler) Server {
	return Server{listener: listener, decoder: decoder, encoder: encoder, handler: handler}
}

// Run слушает и обрабатывает подключения.
func (server *Server) Run() error {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			return err
		}

		go server.runSession(conn)
	}
}

// Stop останавливает работу Server.
func (server *Server) Stop() error {
	err := server.listener.Close()

	server.muForSessions.Lock()
	defer server.muForSessions.Unlock()

	server.isStopped = true

	for _, session := range server.sessions {
		go session.Close()
	}

	return err // [todo] норм ли? прост при ошибке всё равно надо все соединения закрыть будто хз. Поэтому берётся ошибка из самого начала вообще
}

func (server *Server) runSession(conn io.ReadWriteCloser) {
	sessionID := SessionID(server.sessionIDGenerator.NextID())

	server.muForSessions.Lock()
	defer server.muForSessions.Unlock()

	if server.isStopped {
		return
	}

	session := NewSession(sessionID, server.decoder, server.encoder, server.handler, conn)
	session.Run()
	server.sessions[sessionID] = session
	session.SubscribeOnClosure("server", server.sessionsClosures) // [todo] или надо ссылку на канал передавать?
}
