package main

import (
	"fmt"
	"github.com/Oliverwqcwrw/aolifu-go/greetings"
)

type Cat struct {
	Name string
	age  int
}

func main() {
	cat1 := Cat{
		"oliver",
		12,
	}
	cat1.test1()
	fmt.Println(cat1)

	cat1.test2()
	fmt.Println(&cat1)

	sum := cat1.sum(10)
	fmt.Println("sum value is ", sum)

	person1 := greetings.Person{
		Name: "oliver",
		Age:  12,
	}
	fmt.Println(person1)

	student1 := greetings.NewStudent("oliver", 23)
	fmt.Println(*student1)
	fmt.Println(student1.GetAge())

}

func (cat Cat) test1() {
	cat.Name = "bob"
	fmt.Printf("cat name is %v \n", cat.Name)
}

func (cat *Cat) test2() {
	cat.Name = "alice"
	fmt.Printf("cat name is %v \n", cat.Name)
}

func (cat Cat) sum(n int) int {

	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

func (cat *Cat) String() string {
	str := fmt.Sprintf("name = %v, age = %v", cat.Name, cat.age)
	return str
}
