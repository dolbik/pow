package connection

import (
	"encoding/binary"
	"net"

	serviceErrors "github.com/dolbik/pow/pkg/error"
)

type Reader interface {
	Read() ([]byte, error)
}

type IOReader struct {
	conn net.Conn
}

func NewIOReader(conn net.Conn) Reader {
	return &IOReader{
		conn: conn,
	}
}

func (mr *IOReader) Read() ([]byte, error) {
	var body []byte
	var msgSize uint32

	err := binary.Read(mr.conn, binary.BigEndian, &msgSize)

	if err != nil {
		return nil, err
	}
	if msgSize > 4096 {
		return nil, serviceErrors.ErrMessageTooLong
	}
	body = make([]byte, msgSize)
	_, err = mr.conn.Read(body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
