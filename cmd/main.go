package main

import (
	"github.com/go-webserver/cmd/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	server.Start()
}
