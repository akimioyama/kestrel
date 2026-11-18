package ports

import (
	"net"
)

type SessionFactory interface {
	Session(conn net.Conn, router Router) Session
}

type Session interface {
	Run() error
}
