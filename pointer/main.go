package main

import "fmt"

func main() {
	printAddr()
	modifyPointerValue()
}

func printAddr() {
	var i int = 3
	fmt.Println("i address is ", &i)

	var i2 *int = &i
	fmt.Printf("i2 type is %T , i2 value is %v , i2 address is %v \n", i2, i2, &i2)

	fmt.Println("i2 point value is", *i2)
}

func modifyPointerValue() {
	var i int = 4
	var i2 *int = &i
	*i2 = 6
	fmt.Println("i2 point value is ", i)

}
