package main

import (
	"fmt"
	"net"
	"os"

	"github.com/dolbik/pow/pkg/handler"
)

func main() {

	addr, err := handler.ServerAddress()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Cannot start server %s, %v\n", addr, err)
		os.Exit(1)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Cannot close listener: ", err)
		}
	}(listener)

	fmt.Println("Server is starter on ", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Cannot accept incoming connection: ", err)
			os.Exit(1)
		}

		serverHandler := handler.NewServerHandler(conn)
		go serverHandler.Handle()
	}
}
