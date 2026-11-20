package tcp

import (
	"io"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"net"
)

// Server - обработчик подключений.
type Server struct {
	listener net.Listener
	decoder  transport.Decoder
	encoder  transport.Encoder
	handler  ports.Handler
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
	return server.listener.Close() // закрыть все коннекшены???
}

func (server *Server) runSession(conn io.ReadWriteCloser) {
	defer conn.Close()
	session := NewSession(server.decoder, server.encoder, server.handler, conn)
	session.Run()
}
