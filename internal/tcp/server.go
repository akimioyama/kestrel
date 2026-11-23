package tcp

import (
	"io"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"net"
)

// Server - обработчик подключений.
type Server struct {
	listener           net.Listener
	decoder            transport.Decoder
	encoder            transport.Encoder
	handler            ports.Handler
	sessionIDGenerator ports.IDGenerator
	sessionManager     SessionManager
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
	server.sessionManager.CloseAll()
	return err // [todo] норм ли? прост при ошибке всё равно надо все соединения закрыть будто хз. Поэтому берётся ошибка из самого начала вообще
}

func (server *Server) runSession(conn io.ReadWriteCloser) {
	sessionID := SessionID(server.sessionIDGenerator.NextID())
	session := NewSession(sessionID, server.decoder, server.encoder, server.handler, conn)
	session.Run()
	server.sessionManager.AddSession(session)
}
