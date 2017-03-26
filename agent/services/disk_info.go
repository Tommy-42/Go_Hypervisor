package services

import (
	"net"
)

func Disk_Info(conn net.Conn) {
	res := []byte("disk\n")
	_, _ = conn.Write(res)
}
