package transport

import (
	"kestrel/internal/core/broker/models/messages/responses"
	"net"
)

type Writer interface {
	Write(response responses.Response) error
}

type WriterImp struct {
	conn    net.Conn
	encoder Encoder
}

func NewWriter(conn net.Conn, encoder Encoder) WriterImp {
	return WriterImp{conn: conn, encoder: encoder}
}

func (writer *WriterImp) Write(response responses.Response) error {
	encodedResponce, err := writer.encoder.Encode(response)

	if err != nil {
		return nil
	}

	_, err = writer.conn.Write(encodedResponce) //надо ли обрабатывать число записаных байтов? вроде из conn должна ошибка кидаться

	return err
}
