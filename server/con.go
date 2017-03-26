package main

import (
	"fmt"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "4242"
	CONN_TYPE = "tcp"
)

func main() {

	fmt.Println("Entering main")

	conn, ok := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	// Check if there is an error while connecting to the agent
	if ok != nil {
		fmt.Println("Error while contacting the agent")
		return
	}

	// Close the listener when the application closes.
	defer conn.Close()

	cmd := []byte("ping\n")

	bytewrite, ok := conn.Write(cmd)

	fmt.Println("bytewrite: ", bytewrite)

	fmt.Println("Exiting main")
}
