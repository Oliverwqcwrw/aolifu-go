package main

import "fmt"

func main() {
	testChan1()

	traverse()

	channel := make(chan int, 10)
	exitChannel := make(chan bool, 1)
	writeData(channel)
	readData(channel, exitChannel)
	for {
		_, ok := <-exitChannel
		if !ok {
			break
		}
	}

	channel2 := make(chan int, 1000)
	primeChannel := make(chan int, 2000)
	exitChannel2 := make(chan bool, 4)
	go writeData2(channel2)
	for i := 0; i < 4; i++ {
		go primeNum(channel2, primeChannel, exitChannel2)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChannel2
		}
		close(exitChannel2)
		close(primeChannel)
	}()

	for {
		prime, ok := <-primeChannel
		if !ok {
			break
		}
		fmt.Println("prime number is ", prime)
	}
	fmt.Println("main thread is existing")
}

func testChan1() {
	channel1 := make(chan int, 3)
	channel1 <- 2
	channel1 <- 3
	channel1 <- 4

	num := <-channel1
	fmt.Println("num value is", num)

	num2 := <-channel1
	fmt.Println("num2 value is", num2)

	num3 := <-channel1
	fmt.Println(num3)
	close(channel1)

}

func traverse() {
	channel1 := make(chan int, 20)
	for i := 1; i < 20; i++ {
		channel1 <- i * 2
	}
	close(channel1)

	for v := range channel1 {
		fmt.Println(v)
	}
}

func writeData(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
		fmt.Println("write data value is", i)
	}
	close(channel)
}

func readData(channel chan int, exitChannel chan bool) {

	for v := range channel {
		fmt.Println("read data value is", v)
	}
	exitChannel <- true
	close(exitChannel)
}

func writeData2(channel chan int) {
	for i := 2; i < 1000; i++ {
		channel <- i
	}
	close(channel)
}

func primeNum(channel chan int, primeChannel chan int, exitChannel chan bool) {
	for {
		num, ok := <-channel
		if !ok {
			break
		}
		flag := true
		for i := 2; i < num; i++ {
			// it is not prime number
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChannel <- num
		}
	}
	fmt.Println("有一个协程退出了")
	exitChannel <- true
}

func readOnlyOrWriteOnly() {
	// write only
	var channel chan<- int
	channel <- 20

	// read only
	var channel2 <-chan int
	<-channel2

}
