package logic

import (
	"errors"
	"kestrel/internal/core/broker/models/messages/requests"
	"kestrel/internal/core/broker/models/messages/responses"
	"net"
)

type Server struct {
	listener net.Listener
	router   Router
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
	session := NewSession(conn, server.router)
	session.Run()
}

type Reader interface {
	Read() (requests.Request, error)
}
type Writer interface {
	Write(response responses.Response) error
}

type ReaderImp struct {
	decoder Decoder
}

func (reader *ReaderImp) Read() (requests.Request, error) {
	return reader.decoder.Decode() // декодер сразу из коннекшена читает
}

type WriterImp struct {
	conn    net.Conn
	encoder Encoder
}

func (writer *WriterImp) Write(response responses.Response) error {
	encodedResponce, err := writer.encoder.Encode(response)

	if err != nil {
		return nil
	}

	_, err = writer.conn.Write(encodedResponce) //надо ли обрабатывать число записаных байтов? вроде из conn должна ошибка кидаться

	return err
}

type Decoder interface {
	Decode() (requests.Request, error)
}
type Encoder interface {
	Encode(response responses.Response) ([]byte, error)
}

type DecoderImp struct {
	conn net.Conn
}

func (_ *DecoderImp) Decode() (requests.Request, error) {
	return nil, errors.New("not supported")
}

type EncoderImp struct{}

func (_ *EncoderImp) Encode(response responses.Response) ([]byte, error) {
	return nil, errors.New("not supported")
}

type Session struct {
	reader Reader
	writer Writer
	router Router
}

func (session *Session) Run() error {
	for {
		reqest, err := session.reader.Read()
		// может имеет смысл закрывать соединение вместо возвращения ошибки?
		// и вместо ридера врайтера держать де/сериализаторы и коннекшен.
		// но конекшен внутри них тоже будет. Или можно просто тут у ридера и врайтера по стопу вызвать.
		// Или вообще в сервере лямбду вызывать с дефер остановкой. Пока сделал остановкой в сервере
		if err != nil {
			return err
		}

		response := session.router.Route(reqest)

		err = session.writer.Write(response)
		if err != nil {
			return err
		}
	}
}

func NewDecoder(conn net.Conn) Decoder { return &DecoderImp{conn: conn} }
func NewEncoder() Encoder              { return &EncoderImp{} }
func NewReader(conn net.Conn) Reader   { return &ReaderImp{decoder: NewDecoder(conn)} }
func NewWriter(conn net.Conn) Writer   { return &WriterImp{conn: conn, encoder: NewEncoder()} }
func NewSession(conn net.Conn, router Router) Session {
	reader := NewReader(conn)
	writer := NewWriter(conn)
	return Session{reader: reader, writer: writer, router: router}
}
