package main

import "fmt"

func main() {
	var a int = 10 / 4
	fmt.Println("a is ", a)

	var b float32 = 10 / 4
	fmt.Println("b is ", b)

	var c float32 = 10.0 / 4
	fmt.Println("c is ", c)

	mod()
	swap()
}

func mod() {
	// a % b = a - a / b * b
	var a int = 10 % 3
	fmt.Println("a value is ", a)

	var b int = -10 % 3
	fmt.Println("b value is ", b)

	var c int = 10 % -3
	fmt.Println("c value is ", c)

	var d int = -10 % -3
	fmt.Println("d value is ", d)
}

func swap() {
	var a int = 10
	var b int = 20
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a value is %d b value is %d", a, b)
}
