package handler

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/dolbik/pow/pkg/connection"
	"github.com/dolbik/pow/pkg/message"
	"github.com/dolbik/pow/pkg/message/processor"
)

type ClientHandler struct {
	reader connection.Reader
	writer connection.Writer
}

func NewClientHandler(conn net.Conn) *ClientHandler {
	return &ClientHandler{
		reader: connection.NewIOReader(conn),
		writer: connection.NewIOWriter(conn),
	}
}

func (ch *ClientHandler) Handle() {

	writeMessage(ch.writer, processor.ServiceRequest)
	msg := readMessage(ch.reader)

	messageProcessor, err := message.GetMessageProcessor(msg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	resourceMsg, err := messageProcessor.Process(msg)
	writeMessage(ch.writer, resourceMsg.Marshal())
	msg = readMessage(ch.reader)
	fmt.Printf("%s. Client got response: %s\n", time.Now().UTC(), string(msg.Body))
}

func writeMessage(writer connection.Writer, msg []byte) {
	err := writer.Write(msg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func readMessage(reader connection.Reader) *processor.Message {
	response, err := reader.Read()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	msg, err := processor.Parse(response)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return msg
}
