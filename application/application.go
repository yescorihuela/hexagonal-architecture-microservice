package application

import "strconv"

func Run() error {
	serverHost := "localhost"
	serverPort, _ := strconv.ParseUint("8000", 10, 0)
	server := NewServer(serverHost, uint(serverPort))
	return server.Run()
}
