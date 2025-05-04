package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	fmt.Println("🚀 Echo server going up... sit tight!")

	listener, err := net.Listen("tcp", "127.0.0.1:16000")
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()

	done := make(chan struct{})

	for {
		defer func() { done <- struct{}{} }()
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(connection)
	}

}
func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	for {
		_, err := connection.Read(buffer)

		if err != nil {
			if err == io.EOF {
				log.Println("Connection Closed")
			} else {
				log.Println("Error reading from connection: ", err)
			}
			return
		}

		log.Printf("Recieved: %s", string(buffer))
	}

}
