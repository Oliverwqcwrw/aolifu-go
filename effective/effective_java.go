package main

import "fmt"

func main() {
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	person := "Oliver"
	if attended[person] { // 若某人不在此映射中，则为 false
		fmt.Println(person, "正在开会")
	}
	fmt.Println(attended[person])
}
