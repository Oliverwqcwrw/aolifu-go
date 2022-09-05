package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"golang.org/x/example/stringutil"
	"log"
	"oliver.com/greetings"
)

func main() {
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(stringutil.ToUpper("Hello"))

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
}
