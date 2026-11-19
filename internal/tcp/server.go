package tcp

import (
	"io"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"net"
)

type Server struct {
	listener net.Listener
	decoder  transport.Decoder
	encoder  transport.Encoder
	handler  ports.Handler
}

func NewServer(listener net.Listener, decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler) Server {
	return Server{listener: listener, decoder: decoder, encoder: encoder, handler: handler}
}

func (server *Server) Run() error {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			return err
		}

		go server.runSession(conn)
	}
}

func (server *Server) Stop() error {
	return server.listener.Close() // закрыть все коннекшены???
}

func (server *Server) runSession(conn io.ReadWriteCloser) {
	defer conn.Close()
	session := NewSession(server.decoder, server.encoder, server.handler, conn)
	session.Run()
}
