package handler

import (
	"fmt"
	"os"

	serviceErrors "github.com/dolbik/pow/pkg/error"
)

type Handler interface {
	Handle()
}

func ServerAddress() (string, error) {
	serverAddress := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	if serverAddress == "" {
		return "", serviceErrors.ErrServerAddress
	}

	if serverPort == "" {
		return "", serviceErrors.ErrServerPort
	}

	return fmt.Sprintf("%s:%s", serverAddress, serverPort), nil
}
