package main

import (
	"gohypervisor/agent/services"
	"net"
)

var (
	mapService = map[string]func(net.Conn){
		"ping\n": services.Ping,
		// "push_conf": services.PushConf,
	}
)
