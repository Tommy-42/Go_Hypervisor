package services

import (
	"net"
)

func Ping(conn net.Conn) {
	res := []byte("pong\n")
	_, _ = conn.Write(res)
}
