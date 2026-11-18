package transport

import (
	"errors"
	"kestrel/internal/core/broker/models/messages/requests"
	"net"
)

type Decoder interface {
	Decode() (requests.Request, error)
}

type DecoderImp struct {
	conn net.Conn
}

func NewDecoder(conn net.Conn) DecoderImp { return DecoderImp{conn: conn} }

func (_ *DecoderImp) Decode() (requests.Request, error) {
	return nil, errors.New("not supported")
}
