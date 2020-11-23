package main

import (
	"fmt"
	"net"
	"firlus.dev/nononet/internal/connection"
)

func main() {
	ln, err := net.Listen("tcp", ":42002") // TODO Port as env variable
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go connection.HandleConnection(conn)
	}
}