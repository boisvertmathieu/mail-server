package server

import (
	"fmt"
	"log"
	"net"
)

const (
	Host = "localhost"
	Port = "9988"
	Type = "tcp"
)

var clientSocketList []net.Conn
var clientConnectedList []string
var clientConnectedCount int

func Start() {
	for {
		run()
	}
}

func run() {
	fmt.Println("Server starting...")
	server, err := net.Listen(Type, fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the server
	defer func(server net.Listener) {
		err := server.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(server)

	fmt.Println(fmt.Sprintf("Listening on %s:%s", Host, Port))
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Client connected")
		processClient(connection)
	}
}

func processClient(connection net.Conn) {
	// Defer closing the connection
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(connection)

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
	_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
}

func registerNewClient(connection net.Conn, client string) {
	clientConnectedCount++
	_ = append(clientConnectedList, client)
	_ = append(clientSocketList, connection)
}

func unregisterClient(connection net.Conn, client string) {
	clientConnectedCount--
	for i, conn := range clientSocketList {
		if conn == connection {
			_ = append(clientSocketList[:i], clientSocketList[i+1:]...)
		}
	}
	for i, c := range clientConnectedList {
		if c == client {
			_ = append(clientConnectedList[:i], clientConnectedList[i+1:]...)
		}
	}
}
