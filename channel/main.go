package main

import "fmt"

func main() {
	test5()
}

func test5() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("end")
}

func rece(c chan int) {
	x := <-c
	fmt.Println("received ", x)
}

func test4() {
	ch := make(chan int)
	go rece(ch)
	ch <- 10
	fmt.Println("send success")
}

func test3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("send success")
}

func test2() {
	var ch chan int
	fmt.Println(ch)
}

func test1() {
	ch := make(chan int)
	go func() {
		num := 10
		ch <- num
	}()
	receivedNum := <-ch
	fmt.Println(receivedNum)
}
