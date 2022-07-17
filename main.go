package main

import (
	"github.com/boisvertmathieu/mail-server/client"
	"github.com/boisvertmathieu/mail-server/server"
)

func main() {
	server.Start()
	client.Connect()
}
