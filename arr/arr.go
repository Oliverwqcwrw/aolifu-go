package main

import "fmt"

func main() {

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 [3]int = [3]int{4, 5, 6}
	fmt.Println(arr2)

	var arr3 = [...]int{7, 8, 9}
	fmt.Println(arr3)

	var arr4 = [...]int{1: 10, 2: 11, 0: 12}
	fmt.Println(arr4)

	putLetter()

	twoDimensionArr()

	traverseArr()
}

func putLetter() {
	var arr1 = [26]byte{}
	for index, _ := range arr1 {
		arr1[index] = 'A' + byte(index)
	}

	for _, value := range arr1 {
		fmt.Printf("%c ", value)
	}
	fmt.Println()

}

func twoDimensionArr() {
	var arr [4][6]int
	arr[0][2] = 10
	arr[1][3] = 10
	arr[2][4] = 10

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
}

func traverseArr() {
	var arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("value is %v ", arr[i][j])
		}
		fmt.Println()
	}

	for index, _ := range arr {
		for _, value2 := range arr[index] {
			fmt.Printf("value is %v ", value2)
		}
		fmt.Println()
	}
}
