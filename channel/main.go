package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		num := 10
		ch <- num
	}()
	receivedNum := <-ch
	fmt.Println(receivedNum)
}
