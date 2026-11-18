package transport

import "kestrel/internal/core/broker/models/messages/requests"

type Reader interface {
	Read() (requests.Request, error)
}

type ReaderImp struct {
	decoder Decoder
}

func NewReader(decoder Decoder) ReaderImp { return ReaderImp{decoder: decoder} }

func (reader *ReaderImp) Read() (requests.Request, error) {
	return reader.decoder.Decode() // декодер сразу из коннекшена читает
}
