package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	fmt.Println("conn redis succ", conn)
	defer conn.Close()

	TestExpire(conn)
}

func TestExpire(conn redis.Conn) {
	_, err := conn.Do("expire", "a", 10)
	if err != nil {
		fmt.Println("expire failed", err)
		return
	}
}

func TestMultString(conn redis.Conn) {
	_, err := conn.Do("MSet", "a", 10, "b", 20)
	if err != nil {
		fmt.Println("MSet failed", err)
		return
	}

	r, err := redis.Ints(conn.Do("MGet", "a", "b"))
	if err != nil {
		fmt.Println("MGet failed", err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}

}

func TestString(conn redis.Conn) {
	_, err := conn.Do("set", "name", 1)
	if err != nil {
		fmt.Println("set failed", err)
		return
	}

	r, err := redis.Int(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get failed", err)
		return
	}
	fmt.Println(r)

}
