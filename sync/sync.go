package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

var m = make(map[string]int)
var m2 = sync.Map{}

func main() {
	test1()

	var x int64
	atomic.AddInt64(&x, 1)
}

func test1() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			v, _ := m2.Load(key)
			fmt.Printf("k=:%v, v=:%v\n", key, v)
			wg.Done()
		}(i)
	}
	wg.Wait()

}

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}
