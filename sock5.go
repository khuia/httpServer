package main

import (
	"github.com/armon/go-socks5"
)

func socks5Start(c *Config) {

	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port
	if err := server.ListenAndServe("tcp", c.Socks5Addr); err != nil {
		panic(err)
	}

}
