package main

import "fmt"

func main() {

	resp := sum(10, 20)
	fmt.Println("resp value is ", resp)
}

func sum(a int, b int) int {
	defer fmt.Println("defer1 goto")
	defer fmt.Println("defer2 goto")

	return a + b
}
