package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("server run")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("error listening...")
		os.Exit(1)
	}

	defer server.Close()
	fmt.Println("Listening on: ", SERVER_HOST+":"+SERVER_PORT)
	fmt.Println("Waiting for Client")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("error Accepted")
			os.Exit(1)
		}

		go processClient(connection)

	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mlen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("error read...")
	}
	fmt.Println("received: ", string(buffer[:mlen]))

	_, err = connection.Write([]byte("thanks for sent me mail :" + string(buffer[:mlen])))

	defer connection.Close()
}
