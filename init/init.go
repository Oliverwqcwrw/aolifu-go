package main

import "fmt"

var a int = test1()

func main() {

	fmt.Println("main is executed")
	fmt.Println("a value  is ", a)
}

func init() {
	fmt.Println("init is executed")
}

func test1() int {
	fmt.Println("test1 is executed")
	return 10
}
