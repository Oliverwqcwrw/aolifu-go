package main

import (
	"errors"
	"fmt"
)

var errDivByZero = errors.New("division by zero")

func main() {
	Try(func() {
		panic("test panic")
	}, func(err interface{}) {
		fmt.Println(err)
	})

}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errDivByZero
	}
	return x / y, nil
}

func Try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}
