package main

import (
	"errors"
	"fmt"
)

func main() {

	num := 10
	fmt.Printf("num type is %T num value is %v num address is %v \n", num, num, &num)

	var num2 = new(int)
	*num2 = 100
	fmt.Printf("num2 type is %T num2 value is %v num2 address is %v num2 point value is  %v  \n", num2, num2, &num2, *num2)

	test1()
	fmt.Println("execute after test1()")

	//err := test2()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("can not execute")

}

func test1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	num := 10
	num2 := 0
	resp := num / num2
	fmt.Println(resp)
}

func test2() error {
	num := 10
	num2 := 0
	resp := num / num2
	fmt.Println(resp)
	return returnError()
}

func returnError() error {

	return errors.New("cause error")
}
