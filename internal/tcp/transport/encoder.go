package transport

import (
	"errors"
	"kestrel/internal/core/broker/models/messages/responses"
)

type Encoder interface {
	Encode(response responses.Response) ([]byte, error)
}

type EncoderImp struct{}

func NewEncoder() EncoderImp { return EncoderImp{} }

func (_ *EncoderImp) Encode(response responses.Response) ([]byte, error) {
	return nil, errors.New("not supported")
}
