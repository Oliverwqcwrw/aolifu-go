package main

import "fmt"

func main() {
	var person Person
	fmt.Println(person)

	if person.pointer == nil {
		fmt.Println("pointer is nil")
	}
	if person.slice == nil {
		fmt.Println("slice is nil")
	}
	if person.map1 == nil {
		fmt.Println("map1 is nil")
	}

	num1 := 10
	person.pointer = &num1
	fmt.Println("person value is ", person)

	person.slice = make([]int, 2)
	person.slice[0] = 10
	fmt.Println("person value is  ", person)

	person.map1 = make(map[int]int)
	person.map1[2] = 10
	fmt.Println("person value is ", person)
}

type Person struct {
	Name    string
	Age     int
	Address string

	pointer *int
	slice   []int
	map1    map[int]int
}
