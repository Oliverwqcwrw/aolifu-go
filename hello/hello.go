package main

import (
	"fmt"
	"github.com/Oliverwqcwrw/aolifu-go/greetings"
	"github.com/Oliverwqcwrw/aolifu-go/hello/morestrings"
	"github.com/google/go-cmp/cmp"
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
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	for _, message := range messages {
		fmt.Println(message)
	}
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
