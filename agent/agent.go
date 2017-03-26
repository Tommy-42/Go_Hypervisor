package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "4242"
	CONN_TYPE = "tcp"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	msg, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil && len(msg) > 0 {
		fmt.Println("Error Reading: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("CMD RECEIVED: %s", msg)

	fn, found := mapService[msg]
	if !found {
		fmt.Println("NOT found")
		return
	}
	fn(conn)
}

func main() {

	fmt.Println("Agent launching")

	ln, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

	// Check if there is an error while getting a connection on the agent
	if err != nil {
		fmt.Println("Error Accepting: ", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error Accepting: ", err.Error())
			fmt.Println("Agent Exiting")
			os.Exit(1)
		}
		go handleConn(conn)
	}
}
