package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex
var m = make(map[string]int)

func main() {
	test3()
}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func test3() {
	wg.Add(1)
	go hello()
	fmt.Println("main goroutine done")
	wg.Wait()
}

func hello() {
	defer wg.Done()
	fmt.Println("Hello ")
}

func test2() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func write() {
	rwlock.Lock()
	x += 1
	rwlock.Unlock()
	wg.Done()
}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func test1() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}
