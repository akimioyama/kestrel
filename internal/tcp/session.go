package tcp

import (
	"io"
	"kestrel/internal/tcp/ports"
	"kestrel/internal/tcp/transport"
)

// Session - сессия работы с подключением.
type Session struct {
	decoder    transport.Decoder
	encoder    transport.Encoder
	handler    ports.Handler
	readWriter io.ReadWriter
}

// NewSession создаёт новую Session.
func NewSession(decoder transport.Decoder, encoder transport.Encoder, handler ports.Handler, readWriter io.ReadWriter) *Session {
	return &Session{decoder: decoder, encoder: encoder, handler: handler, readWriter: readWriter}
}

// Run запускает сессию.
func (session *Session) Run() error {
	for {
		reqest, err := session.decoder.Decode(session.readWriter)

		if err != nil {
			return err
		}

		response := session.handler.Handle(reqest)

		encodedResponce, err := session.encoder.Encode(response)

		if err != nil {
			return err
		}

		_, err = session.readWriter.Write(encodedResponce)

		if err != nil {
			return err
		}
	}
}
