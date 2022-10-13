package main

import (
	"encoding/json"
	"fmt"
)

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

	jsonPerson()
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`

	pointer *int
	slice   []int
	map1    map[int]int
}

func definePerson() {
	var person1 Person
	person1.Name = "a"

	person2 := new(Person)
	person2.Name = "b"

	person3 := &Person{}
	person3.Name = "c"

	person4 := Person{}
	person4.Name = "d"
}

func jsonPerson() {
	person1 := Person{}
	person1.Name = "oliver"
	person1.Address = "beiJing"
	person1.Age = 12

	jsonStr, err := json.Marshal(person1)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("person1 json str is ", string(jsonStr))
}
