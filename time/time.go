package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Printf("now type is %T now value is %v \n", now, now)

	str := fmt.Sprintf("now value is %d-%d-%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(),
		now.Minute(), now.Second())
	fmt.Println(str)

	fmt.Println("year is ", now.Year())
	fmt.Println("month is ", now.Month())
	fmt.Println("day is ", now.Day())
	fmt.Println("hour is ", now.Hour())
	fmt.Println("minute is ", now.Minute())
	fmt.Println("second is ", now.Second())

	str2 := fmt.Sprintf(now.Format("2006-01-02 15:04:05"))
	fmt.Println(str2)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
}

func sleep() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond)
	}
}
