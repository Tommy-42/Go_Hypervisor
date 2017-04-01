package main

import (
	"gosupervisor/agent/services"
	"net"
)

var (
	mapService = map[string]func(net.Conn){
		"ping\n":      services.Ping,
		"disk_info\n": services.Disk_Info,
	}
)
