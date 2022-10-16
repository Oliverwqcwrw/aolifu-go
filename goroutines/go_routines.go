package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	go test2()
	test1()
	fmt.Println(runtime.NumCPU())
	factorial()

}

func test1() {
	c := make(chan int) // 分配一个信道
	// 在Go程中启动排序。当它完成后，在信道上发送信号。
	go func() {
		fmt.Println("async execute")
		c <- 1 // 发送信号，什么值无所谓。
	}()
	fmt.Println("test goroutines")
	<-c // 等待排序结束，丢弃发来的值。
}

func test2() {
	fmt.Println("test2 is  executing")
}

var factorialMap map[int]int = make(map[int]int)
var lock sync.Mutex

func factorial() {
	for i := 2; i < 20; i++ {
		go calFactorial(i)
	}
	time.Sleep(time.Second * 5)
	lock.Lock()
	for k, v := range factorialMap {
		fmt.Printf("map[%v] = %v \n", k, v)
	}
	lock.Unlock()
}

func calFactorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	lock.Lock()
	factorialMap[n] = result
	lock.Unlock()
}
