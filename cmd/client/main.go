package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/dolbik/pow/pkg/handler"
)

func main() {

	addr, err := handler.ServerAddress()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", addr)
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Cannot close connection: %v", err)
		}
	}(conn)
	if err != nil {
		fmt.Printf("Cannot connect to server %s, %v", addr, err)
		os.Exit(1)
	}

	clientHandler := handler.NewClientHandler(conn)

	for {
		clientHandler.Handle()
		time.Sleep(time.Second * 10)
	}
}
