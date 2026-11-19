package tcp

import (
	"io"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
)

type SessionImp struct {
	decoder    transport.Decoder
	encoder    transport.Encoder
	handler    ports.Handler
	readWriter io.ReadWriter
}

func NewSession(decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler, readWriter io.ReadWriter) SessionImp {
	return SessionImp{decoder: decoder, encoder: encoder, handler: handler, readWriter: readWriter}
}

func (session *SessionImp) Run() error {
	for {
		reqest, err := session.decoder.Decode(session.readWriter)

		if err != nil {
			return err
		}

		response := session.handler.Handle(reqest)

		encodedResponce, err := session.encoder.Encode(response)

		if err != nil {
			return nil
		}

		_, err = session.readWriter.Write(encodedResponce)

		if err != nil {
			return err
		}
	}
}
