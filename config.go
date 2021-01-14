package main

type Config struct {
	Domain string
	IPAddr string
	Port   int
}

var MyConfig = Config{
	Domain: "auth.npool.com",
	IPAddr: "106.14.125.55",
	Port:   40001,
}
