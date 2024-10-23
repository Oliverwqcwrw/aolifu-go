package main

import (
	"fmt"
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
	var age int
	fmt.Scanf("%d", &age)
	fmt.Println("You are", age, "years old!")
	if count := 10; count > 5 {
		fmt.Println("Count is greater than 5")
	} else {
		fmt.Println("Count is not greater than 5")
	}
}
