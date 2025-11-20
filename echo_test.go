package main

import (
	"net"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	go startTCPEchoServer("2000")

	time.Sleep(200 * time.Millisecond)

	connection, err := net.Dial("tcp", "127.0.0.1:2000")

	if err != nil {
		t.Fatal(err)
	}

	defer connection.Close()

	test := "Hello world!"

	expect := "Hello world!"

	connection.Write([]byte(test))
	data := make([]byte, 50)
	n, error := connection.Read(data)

	if error != nil {
		t.Fatal(error)
	}

	if string(data[:n]) != expect {
		t.Fatalf("Was expected %s, got %s", expect, string(data))
	}
}
