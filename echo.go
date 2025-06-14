package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func main() {
	// Get protocol from command line flag, default to tcp
	protocol := flag.String("proto", "tcp", "protocol to use: tcp or udp")
	port := flag.String("port", "7", "port number to listen on")
	flag.Parse()

	log.Printf("🚀 Echo server using protocol %s on port %s...\n", *protocol, *port)

	if *protocol == "tcp" {
		startTCPEchoServer(*port)
	} else if *protocol == "udp" {
		startUDPEchoServer(*port)
	} else {
		log.Fatalf("Unknown protocol: %s. Use tcp or udp.", *protocol)
	}
}

func startUDPEchoServer(port string) {
	s, err := net.ListenPacket("udp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := s.ReadFrom(buf)
		if err != nil {
			log.Println("Error reading udp packet:", err)
			continue
		}

		_, err = s.WriteTo(buf[:n], addr)
		if err != nil {
			log.Println(err)
		}
	}
}

func startTCPEchoServer(port string) {

	listener, err := net.Listen("tcp", ":"+port)

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
					log.Println("FIN packet for: ", connection.RemoteAddr(), " Closing the tcp session...")
					return
				}
				connection.Write(data[:n])
			}
		}(connection)
	}
}
