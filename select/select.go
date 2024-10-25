package main

import (
	"fmt"
	"time"
)

func main() {
	test2()
}

func test1() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

func test2() {
	var resChan = make(chan int)
	go func() {
		resChan <- 1
	}()
	select {
	case data := <-resChan:
		fmt.Println("data value is ", data)
	case <-time.After(3 * time.Second):
		fmt.Println("request timeout")
	}

}
