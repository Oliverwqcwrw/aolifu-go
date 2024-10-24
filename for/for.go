package main

import "fmt"

func main() {
	sum := commonLoop(10, 20, 30)
	fmt.Println("sum value is ", sum)

	rangeLoop()
	rangeStr()

}

func commonLoop(a int, args ...int) int {
	sum := a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func rangeLoop() {
	var arr1 = [3]int{1, 2, 3}
	for _, value := range arr1 {
		fmt.Println(value)
	}
}

func rangeStr() {
	str := "hello"
	for _, value := range str {
		fmt.Println(value)
	}
}
