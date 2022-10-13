package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Animal interface {
	Road()
}

type Person interface {
	Animal
	Study()
	Walk()
}

type People struct {
}

func (people People) Study() {
	fmt.Println("people start to study")
}

func (people People) Walk() {
	fmt.Println("people start  to walk")
}

func (people People) Road() {
	fmt.Println("people start road")
}

type Student struct {
}

func (student Student) Study() {
	fmt.Println("student start to study")
}

func (student Student) Walk() {
	fmt.Println("student start to walk")
}

func (student Student) Road() {
	fmt.Println("student start to road")
}

type AdultSlice []Adult

type Adult struct {
	name string
	age  int
}

func (adult Adult) action(person Person) {
	person.Study()
	person.Walk()
	person.Road()
}

func (adultSlice AdultSlice) Len() int {
	return len(adultSlice)
}

func (adultSlice AdultSlice) Less(i, j int) bool {
	return adultSlice[i].age < adultSlice[j].age
}

func (adultSlice AdultSlice) Swap(i, j int) {
	//temp := adultSlice[i]
	//adultSlice[i] = adultSlice[j]
	//adultSlice[j] = temp
	adultSlice[i], adultSlice[j] = adultSlice[j], adultSlice[i]
}

func sortAdult() {
	var slice AdultSlice
	for i := 0; i < 10; i++ {
		adult := Adult{name: fmt.Sprintf("oliver %d", rand.Intn(100)),
			age: rand.Intn(100)}
		slice = append(slice, adult)
	}
	fmt.Println(slice)

	sort.Sort(slice)

	fmt.Println(slice)
}

func main() {
	adult := Adult{}
	student := Student{}
	people := People{}
	adult.action(student)
	adult.action(people)

	sortAdult()
}
