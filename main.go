package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yescorihuela/agrak/application"
)

func main() {
	if err := application.Run(); err != nil {
		log.WithError(err).Fatalln("Fatal error")
	}
}
