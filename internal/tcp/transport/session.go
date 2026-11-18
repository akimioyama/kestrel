package transport

import "kestrel/internal/tcp/ports"

type SessionImp struct {
	reader Reader
	writer Writer
	router ports.Router
}

func NewSession(reader Reader, writer Writer, router ports.Router) SessionImp {
	return SessionImp{reader: reader, writer: writer, router: router}
}

func (session *SessionImp) Run() error {
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
