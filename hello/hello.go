package main

import (
	"fmt"
	"github.com/Oliverwqcwrw/rocketmq-client-go/greetings"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/example/stringutil"
	"log"
	"os"
)

func init() {
	var (
		home   = os.Getenv("HOME")
		user   = os.Getenv("USER")
		gopath = os.Getenv("GOPATH")
	)
	fmt.Println(home, user, gopath)

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
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(stringutil.Reverse("Hello"))

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Oliver")
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	//message := greetings.Hello("Gladys")
	fmt.Println(messages)

	x := []int{1, 2, 3}
	x = append(x, 4, 5, 6)
	fmt.Println(x)

	x = []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
	log.Println("Stats In One Minute.", map[string]interface{}{
		"statsName": "statsName",
		"statsKey":  "statsKey",
		"SUM":       "SUM",
		"TPS":       fmt.Sprintf("%.2f", 2.3),
		"AVGPT":     23,
	})

}
