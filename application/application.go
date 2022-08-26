package application

import "strconv"

func Run() error {
	serverHost := "localhost"                         // This value would be defined by envvar
	serverPort, _ := strconv.ParseUint("8000", 10, 0) // This value would be defined by envvar
	server := NewServer(serverHost, uint(serverPort))
	return server.Run()
}
