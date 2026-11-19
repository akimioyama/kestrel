package transport

import (
	"errors"
	"io"
	"kestrel/internal/core/broker/models/messages/requests"
)

type Decoder interface {
	Decode(reader io.Reader) (requests.Request, error)
}

type DecoderImp struct{}

func NewDecoder() DecoderImp { return DecoderImp{} }

func (_ *DecoderImp) Decode(reader io.Reader) (requests.Request, error) {
	return nil, errors.New("not supported")
}
