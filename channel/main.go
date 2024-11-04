package main

import (
	"fmt"
	"math/rand"
)

func main() {
	test8()
}

func test8() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	i := 0
	for {
		if i > 300 {
			break
		}
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
		i++
	}
}

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func test7() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func test6() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			if data, ok := <-ch1; ok {
				ch2 <- data + 2
			} else {
				break
			}
		}
		close(ch2)
	}()

	for i := range ch2 {
		fmt.Println(i)
	}

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
