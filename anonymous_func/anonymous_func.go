package main

import (
	"fmt"
	"strings"
)

var (
	testFunc2 = func(a int, b int) int {
		return a * b
	}
)

func main() {

	resp := func(a int, b int) int {
		return a + b
	}(10, 20)

	fmt.Println("resp value is ", resp)

	testFunc := func(a int, b int) int {
		return a - b
	}

	resp2 := testFunc(20, 10)
	fmt.Println("resp2 value is ", resp2)

	resp3 := testFunc2(10, 20)
	fmt.Println("resp3 value is ", resp3)

	testFunc3 := accumulator()
	resp4 := testFunc3(10)
	fmt.Println("resp4 value is ", resp4)

	resp5 := testFunc3(10)
	fmt.Println("resp5 value is ", resp5)

	testFunc4 := formatSuffix(".jpg")
	fileName := testFunc4("bird")
	fmt.Println("fileName is ", fileName)

}

func accumulator() func(a int) int {
	a := 10
	return func(num int) int {
		a += num
		return a
	}
}

// 使用闭包 可以减少传递参数的次数，可以把引用到的变量保存下来
func formatSuffix(suffix string) func(fileName string) string {
	return func(fileName string) string {
		if strings.HasSuffix(fileName, suffix) {
			return fileName
		} else {
			return fileName + ".jpg"
		}
	}
}
