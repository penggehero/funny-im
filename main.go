package main

import log "github.com/sirupsen/logrus"

func main() {
	server := newIMServer()
	if err := server.Run(); err != nil {
		log.Fatalf("server run error: %v", err)
	}
}
