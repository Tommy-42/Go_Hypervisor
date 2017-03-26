package services

import (
	"fmt"
	"net"
)

func Ping(conn net.Conn) {
	fmt.Println("pong")
}

func PushConf(s string) {}
