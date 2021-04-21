package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	var data []byte
	d, e := conn.Read(data)
	fmt.Printf("%s | %v | %v | %s \n", conn.RemoteAddr(), d, e, data)
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}