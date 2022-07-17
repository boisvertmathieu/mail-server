package client

import (
	"fmt"
	"github.com/boisvertmathieu/mail-server/server"
	"log"
	"net"
)

func Connect() {
	connection, err := net.Dial(server.Type, fmt.Sprintf("%s:%s", server.Host, server.Port))
	if err != nil {
		log.Fatal(err)
	}

	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(connection)

	_, err = connection.Write([]byte("Hello Server! Greetings."))
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	} else {
		fmt.Println("Received: ", string(buffer[:mLen]))
	}
}
