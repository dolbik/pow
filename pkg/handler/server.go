package handler

import (
	"fmt"
	"net"
	"time"

	"github.com/dolbik/pow/pkg/connection"
	"github.com/dolbik/pow/pkg/message"
	"github.com/dolbik/pow/pkg/message/processor"
)

type ServerHandler struct {
	conn   net.Conn
	reader connection.Reader
	writer connection.Writer
}

func NewServerHandler(conn net.Conn) *ServerHandler {
	return &ServerHandler{
		conn:   conn,
		reader: connection.NewIOReader(conn),
		writer: connection.NewIOWriter(conn),
	}
}

func (sh *ServerHandler) Handle() {
	defer func(conn net.Conn) {
		err := sh.conn.Close()
		if err != nil {
			sh.error(err)
		}
	}(sh.conn)

	for {
		data, err := sh.reader.Read()

		if err != nil {
			sh.error(err)
			return
		}

		fmt.Printf(
			"%s. Server got request from %s: \"%s\"\n",
			time.Now().UTC(),
			sh.conn.RemoteAddr().String(),
			string(data),
		)

		msg, err := processor.Parse(data)
		if err != nil {
			sh.error(err)
			return
		}

		msg.Resource = sh.conn.RemoteAddr().String()

		messageProcessor, err := message.GetMessageProcessor(msg)
		if err != nil {
			sh.error(err)
			return
		}

		response, err := messageProcessor.Process(msg)
		if err != nil {
			sh.error(err)
			return
		}

		err = sh.writer.Write(response.Marshal())
		if err != nil {
			sh.error(err)
			return
		}

		fmt.Printf(
			"%s. Server sent response to %s: \"%s\"\n",
			time.Now().UTC(),
			sh.conn.RemoteAddr().String(),
			string(response.Marshal()),
		)
	}
}

func (sh *ServerHandler) error(err error) {
	fmt.Printf(
		"%s. Server sent error to %s: %s\n",
		time.Now().UTC(),
		sh.conn.RemoteAddr().String(),
		err,
	)
	_ = sh.writer.Write([]byte(err.Error()))
}
