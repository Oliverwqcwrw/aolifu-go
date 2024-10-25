package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	map1 := make(map[int]Student)
	map1[1] = Student{id: 1, name: "bob", age: 12}
	map1[2] = Student{id: 2, name: "bob2", age: 13}
	fmt.Println(map1)
	delete(map1, 1)
	fmt.Println(map1)
}

type User struct {
	string
	int
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`

	pointer *int
	slice   []int
	map1    map[int]int
}

func (p *Person) setAge(newAge int) {
	p.Age = newAge
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

type Animal struct {
	Name string
}

func (a *Animal) move() {
	fmt.Printf("%s 会动 \n", a.Name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s 会wang", d.Name)
}

type Student struct {
	id   int
	name string
	age  int
}
