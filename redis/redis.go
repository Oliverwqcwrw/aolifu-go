package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()
	pool.Close()

	TestHash(conn)
}

func TestHash(conn redis.Conn) {
	_, err := conn.Do("hset", "books", "name1", "a")
	if err != nil {
		fmt.Println("hset failed", err)
		return
	}
	str, err := redis.String(conn.Do("hget", "books", "name1"))
	if err != nil {
		fmt.Println("hget failed", err)
		return
	}
	fmt.Println("hget succ", str)
}

func TestList(conn redis.Conn) {
	_, err := conn.Do("lpush", "name_list", "a", "b", "c")
	if err != nil {
		fmt.Println("lpush failed", err)
		return
	}
	conn.Do("lpop", "name_list")
	str, err := redis.String(conn.Do("lpop", "name_list"))
	if err != nil {
		fmt.Println("lpop failed", err)
	}
	fmt.Println("pop succ", str)

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
