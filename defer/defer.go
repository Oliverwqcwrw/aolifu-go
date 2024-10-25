package main

import "fmt"

func main() {

	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

func sum(a int, b int) int {
	defer fmt.Println("defer1 goto")
	defer fmt.Println("defer2 goto")

	return a + b
}
