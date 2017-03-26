package main

import (
	"fmt"
	"net"
	"time"
)

const (
	//	CONN_HOST = "172.16.17.128"
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "4242"
	CONN_TYPE = "tcp"
)

func handleRequest(conn net.Conn, v string) {

	// Close the listener when the application closes.
	defer conn.Close()

	_, _ = conn.Write([]byte(v))

}

func main() {

	fmt.Println("Entering main")

	for _, v := range services {

		conn, ok := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)

		// Check if there is an error while connecting to the agent
		if ok != nil {
			fmt.Println("Error while contacting the agent")
			return
		}
		go handleRequest(conn, v)
	}

	// bytewrite, ok := conn.Write(cmd)

	//fmt.Println("bytewrite: ", bytewrite)
	time.Sleep(time.Minute * 2)
	fmt.Println("Exiting main")
}
