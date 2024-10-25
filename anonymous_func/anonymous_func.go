package main

import (
	"strings"
)

var (
	testFunc2 = func(a int, b int) int {
		return a * b
	}
)

func main() {
	// --- channel of function ---
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello, World!" }
	println((<-fc)())
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
