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
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

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

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (self *User) Test() {
	fmt.Printf("%p, %v\n", self, self)
}

type Person struct {
	name string
	sex  string
	age  int
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
