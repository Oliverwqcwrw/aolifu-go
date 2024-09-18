package main

import (
	"fmt"
	"github.com/Oliverwqcwrw/aolifu-go/greetings"
	"golang.org/x/example/hello/reverse"
	"log"
	"os"
)

func init() {
	home := os.Getenv("HOME")
	user := os.Getenv("USER")
	gopath := os.Getenv("GOPATH")

	if user == "" {
		log.Fatal("$USER not set")
	}
	if home == "" {
		home = "/home/" + user
	}
	if gopath == "" {
		gopath = home + "/go"
	}

}

func main() {
	fmt.Println(reverse.String("Hello"))
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
