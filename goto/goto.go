package main

import "fmt"

func main() {

	fmt.Println("Hello1")
	goto LABEL1
	fmt.Println("Hello2")
	fmt.Println("Hello3")
LABEL1:
	fmt.Println("Hello4")
	fmt.Println("Hello5")

}
