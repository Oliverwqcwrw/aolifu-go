package main

import (
	"fmt"
)

type People struct {
	Name string
}

type Person struct {
	Name string
}

func getType(obj ...interface{}) {
	for i, v := range obj {
		switch v.(type) {
		case bool:
			fmt.Printf("第 %v 个参数 type is %v", i, "bool \n")
		case int:
			fmt.Printf("第 %v 个参数 type is %v \n", i, "int")
		case Person:
			fmt.Printf("第 %v 个参数 type is %v \n", i, "Person")
		case *Person:
			fmt.Printf("第 %v 个参数 type is %v \n", i, "*Person")
		default:
		}
	}
}

func main() {
	var a interface{}
	a = People{
		Name: "oliver",
	}
	if people2, ok := a.(People); ok {
		fmt.Println(people2)
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println("continue executing")

	num1 := true
	num2 := 1
	num3 := Person{
		Name: "oliver",
	}
	getType(num1, num2, &num3)

}
