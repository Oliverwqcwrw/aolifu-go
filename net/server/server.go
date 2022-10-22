package main

import (
	"fmt"
	"net"
)

func main() {

	lister, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Println("listen failed", err)
	}
	for {
		conn, err := lister.Accept()
		fmt.Println("accept connection")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("conn is", conn, "remoteAddress is", conn.RemoteAddr().String())
		}
		go processContent(conn)
	}
}

func processContent(conn net.Conn) {
	defer conn.Close()
	for {
		byteSlice := make([]byte, 1024)
		n, err := conn.Read(byteSlice)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("accept content is", string(byteSlice[:n]))
	}
}
