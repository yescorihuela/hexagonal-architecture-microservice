package application

import (
	"os"
	"strconv"
)

func Run() error {
	serverHost := os.Getenv("BACKEND_IP")
	serverPort, _ := strconv.ParseUint(os.Getenv("BACKEND_PORT"), 10, 0) // This value would be defined by envvar
	server := NewServer(serverHost, uint(serverPort))
	return server.Run()
}
