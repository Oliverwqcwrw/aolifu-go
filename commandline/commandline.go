package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flagParse()

}

func osArgs() {
	fmt.Println("command line args length is  ", len(os.Args))

	for i, v := range os.Args {
		fmt.Printf("os[%v]=%v \t\n", i, v)
	}
}

func flagParse() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "the default value of user is nil")
	flag.StringVar(&pwd, "pwd", "", "default value of pwd is nil")
	flag.StringVar(&host, "h", "127.0.0.1", "default value of host is 127.0.0.1")
	flag.IntVar(&port, "p", 8080, "default value of port is 8080")

	flag.Parse()
	fmt.Printf("user is %v, pwd is %v, host is %v, port is %v", user, pwd, host, port)
}
