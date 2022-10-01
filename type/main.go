package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var c1 int = 90
	fmt.Printf("c1 = %c \n", c1)

	var c2 byte = 89
	fmt.Printf("c2 = %c\n", c2)

	var c3 byte = 'c'
	fmt.Printf("c3 = %c\n", c3)

	var b1 bool = false
	fmt.Println("b1 = ", b1)
	fmt.Println("b1 size :", unsafe.Sizeof(b1))

	var s1 string = "中国"
	fmt.Println(s1)

	var s2 string = `aaaa\naaaa\t\\\\aaa`
	fmt.Println(s2)

	var s3 string = "aaa" +
		"bbb" + "ccc" +
		"ddd" + "eee"
	fmt.Println(s3)

	var int0 int
	var string0 string
	var bool0 bool
	var float0 float64
	fmt.Println(int0, string0, bool0, float0)

	var int1 int32 = 100
	var float1 float32 = float32(int1)
	fmt.Println(float1)
}
