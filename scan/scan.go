package main

import "fmt"

func main() {
	var name string

	fmt.Println("请输入姓名:")
	// fmt.Scanln(&name)

	//fmt.Printf("name value  is %v", name)

	fmt.Scanf("%s", &name)
	fmt.Println("name value is ", name)

}
