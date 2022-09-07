package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // 分配一个信道
	// 在Go程中启动排序。当它完成后，在信道上发送信号。
	go func() {
		fmt.Println("async execute")
		c <- 1 // 发送信号，什么值无所谓。
	}()
	fmt.Println("test goroutines")
	<-c // 等待排序结束，丢弃发来的值。
}
