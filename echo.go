package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.Println("🚀 Echo server firing up...")

	listener, err := net.Listen("tcp", ":7")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on: ", listener.Addr().String())

	for {
		connection, err := listener.Accept()
		log.Println("New Connection: ", connection.RemoteAddr())
		if err != nil {
			log.Println(err.Error())
		}
		go func(connection net.Conn) {

			defer connection.Close()
			for {
				data := make([]byte, 1024)

				n, err := connection.Read(data)

				if err != nil {
					if err != io.EOF {
						log.Print(err)
					}
					log.Println("FIN packet for: ", connection.RemoteAddr()," Closing the tcp session...")
					return
				}
				connection.Write(data[:n])
			}
		}(connection)
	}
}
