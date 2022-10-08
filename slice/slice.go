package main

import "fmt"

func main() {

	var arr1 = [4]int{1, 2, 3, 4}
	slice1 := arr1[1:3]
	fmt.Printf("slice1 value is %v, slice1 len is %v, slice1 cap is %v \n", slice1, len(slice1), cap(slice1))

	makeSlice()

	makeSlice2()

	traverse()

	slice2 := arr1[:]
	fmt.Println(slice2)

	slice3 := arr1[0:]
	fmt.Println(slice3)

	slice4 := arr1[:3]
	fmt.Println(slice4)

	appendItem()

	copySlice()

	strSlice()
}

func makeSlice() {
	var arr1 []int = make([]int, 4, 8)
	arr1[0] = 10
	arr1[1] = 20
	fmt.Printf("arr1 value is %v, arr1 len is %v, arr1 capacity is %v \n", arr1, len(arr1), cap(arr1))
}

func makeSlice2() {
	var arr1 []int = []int{1, 2, 3}
	fmt.Printf("arr1 value is %v, arr1 len is %v, arr1 capacity is %v \n", arr1, len(arr1), cap(arr1))
}

func traverse() {
	var arr1 = [...]int{1, 2, 3, 4}
	slice1 := arr1[1:4]
	for index, value := range slice1 {
		fmt.Printf("index =  %v value = %v \n", index, value)
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("index = %v value = %v \n", i, slice1[i])
	}
}

func appendItem() {
	var slice1 = []int{1, 2, 3}
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(slice1)
}

func copySlice() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 5, 10)
	copy(slice2, slice1)
	fmt.Println("slice1 value  is ", slice1)
	fmt.Println("slice2 value is ", slice2)
}

func strSlice() {
	str1 := "hello World"
	slice1 := str1[3:]
	fmt.Println(slice1)

}
