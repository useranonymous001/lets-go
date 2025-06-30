package server

import (
	"fmt"
	"log"
	"net"
)

func ListenIncomingRequest(addr string) {

	// addr => "localhost:2000" || ":2000"

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Invalid Addr: ", err)
	}

	fmt.Println("Listening on addr", addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error connecting:", err)
		}

		go handleConnection(conn)

	}

}
