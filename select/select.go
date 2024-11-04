package main

import (
	"fmt"
	"time"
)

func main() {
	test7()
}

func test7() {
	output1 := make(chan string, 10)
	go write(output1)
	for s := range output1 {
		fmt.Println("read", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel is full")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func test6() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	go func() {
		ch1 <- "test1"
	}()
	go func() {
		ch2 <- "hello"
	}()
	select {
	case i := <-ch1:
		fmt.Println("i = ", i)
	case s := <-ch2:
		fmt.Println("s = ", s)
	}
	fmt.Println("main...over...")
}

func test5() {
	// 创建管道
	ch3 := make(chan string)
	ch4 := make(chan string)
	// 启动协程，向管道写入数据
	go test3(ch3)
	go test4(ch4)
	// select 读取每个管道的数据
	select {
	case s3 := <-ch3:
		fmt.Println("s3 = ", s3)
	case s4 := <-ch4:
		fmt.Println("s4 = ", s4)
	}

}

func test3(ch chan string) {
	time.Sleep(5 * time.Second)
	ch <- "test3"
}

func test4(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "test4"
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
