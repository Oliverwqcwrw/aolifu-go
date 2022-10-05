package main

import (
	"fmt"
	"strconv"
)

func main() {

	intToStr()
	floatToStr()
	boolToStr()
	strToInt()
	strToFloat()
	strToBool()

}

func intToStr() {
	var a int = 3
	str := fmt.Sprintf("%d", a)
	fmt.Printf("str type is %T \n", str)

	str2 := strconv.FormatInt(int64(a), 10)
	fmt.Printf("str2 type is %T, str2 = %v \n", str2, str2)

	str3 := strconv.Itoa(a)
	fmt.Printf("str3 type is %T, str3 = %v \n", str3, str3)
}

func floatToStr() {
	var a float32 = 3.10
	str := fmt.Sprintf("%f", a)
	fmt.Printf("str type is %T, str=%v \n", str, str)

	str2 := strconv.FormatFloat(float64(a), 'f', 10, 32)
	fmt.Printf("str2 type is %T, str2= %v \n", str2, str2)
}

func boolToStr() {
	var a bool = true
	str := fmt.Sprintf("%t", a)
	fmt.Printf("str type is %T, str=%v \n", str, str)

	str2 := strconv.FormatBool(a)
	fmt.Printf("str2 type is %T, str2 = %v \n", str2, str2)
}

func strToInt() {
	var str string = "123"
	a, error := strconv.ParseInt(str, 10, 32)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Printf("a type is %T, a value is %v \n", a, a)
}

func strToFloat() {
	var str string = "2.01"
	float1, error := strconv.ParseFloat(str, 5)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Printf("float1 type is %T, float1 value is %v \n", float1, float1)
}

func strToBool() {
	var str string = "true"
	bool1, error := strconv.ParseBool(str)
	if error != nil {
		fmt.Println(error.Error())
	}
	fmt.Printf("bool1 type is %T, bool1 value is %v \n", bool1, bool1)
}
