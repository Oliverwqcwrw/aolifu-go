package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connection is ", conn, "localAddress is", conn.LocalAddr().String())

	reader := bufio.NewReader(os.Stdin)
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	n, err := conn.Write([]byte(content))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("client send %d byte data", n)
}
