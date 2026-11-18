package tcp

import (
	"kestrel/internal/tcp/ports"
	"net"
)

type Server struct {
	listener       net.Listener
	router         ports.Router
	sessionFactory ports.SessionFactory
}

func (server *Server) Run() error {
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			return err
		}

		go server.RunSession(conn)
	}
}

func (server *Server) Stop() error {
	return server.listener.Close()
}

func (server *Server) RunSession(conn net.Conn) {
	defer conn.Close()
	session := server.sessionFactory.Session(conn, server.router)
	session.Run()
}
