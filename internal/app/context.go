package app

import (
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
	"net"
)

type Context struct{}

func (c *Context) Session(conn net.Conn, router ports.Router) ports.Session { // реализация tcp.ports.SessionFactory. А тут, т.к. это часть DI будто.
	decoder := transport.NewDecoder(conn)
	reader := transport.NewReader(&decoder)
	encoder := transport.NewEncoder()
	writer := transport.NewWriter(conn, &encoder)
	session := transport.NewSession(&reader, &writer, router)
	return &session
}
