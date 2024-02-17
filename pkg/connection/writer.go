package connection

import (
	"encoding/binary"
	"net"
)

type Writer interface {
	Write(msg []byte) error
}

type IOWriter struct {
	conn net.Conn
}

func NewIOWriter(conn net.Conn) Writer {
	return &IOWriter{
		conn: conn,
	}
}

func (mw *IOWriter) Write(msg []byte) error {
	err := binary.Write(mw.conn, binary.BigEndian, uint32(len(msg)))
	if err != nil {
		return err
	}
	_, err = mw.conn.Write(msg)
	if err != nil {
		return err
	}

	return nil
}
