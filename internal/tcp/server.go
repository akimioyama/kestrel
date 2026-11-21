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
	sessionIdGenerator ports.IDGenerator
	sessionsClosures   chan *Session
	sessions           map[SessionId]*Session
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

	return err
}

func (server *Server) runSession(conn io.ReadWriteCloser) {
	sessionId := SessionId(server.sessionIdGenerator.NextID())

	server.muForSessions.Lock()
	defer server.muForSessions.Unlock()

	if server.isStopped {
		return
	}

	session := NewSession(sessionId, server.decoder, server.encoder, server.handler, conn)
	session.Run()
	server.sessions[sessionId] = session
	session.SubscribeOnClosure("server", server.sessionsClosures) // [todo] или надо ссылку на канал передавать?
}
