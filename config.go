package main

import (
	"net"
)

type Config struct {
	Domain        string
	IPAddr        *net.IPAddr
	Port          int
	EtcdEndPoints []string
}

const MyConfig = Config{
	Domain:        "auth.npool.com",
	IPAddr:        net.ResolveIPAddr("ip", "106.14.125.55"),
	Port:          40001,
	EtcdEndPoints: []string{"106.14.125.55:2379"},
}
