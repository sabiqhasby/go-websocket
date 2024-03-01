package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error dial...")

	}
	defer connection.Close()

	_, err = connection.Write([]byte("hello server! jubaidah"))
	if err != nil {
		fmt.Println("connection error")
	}

	//READ MESSAGE FROM SERVER
	buffer := make([]byte, 1024)
	mlen, err := connection.Read(buffer)

	if err != nil {
		fmt.Println("Error read")
	}
	fmt.Println("Received :", string(buffer[:mlen]))
}
