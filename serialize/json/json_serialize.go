package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func main() {
	serializeStruct()

	serializeMap()

	serializeSlice()
}

func serializeStruct() {
	person := Person{
		Name:  "oliver",
		Age:   12,
		Phone: "13611111111",
	}
	data, err := json.Marshal(&person)
	if err != nil {
		fmt.Println("parse error")
		return
	}
	fmt.Println("json serialize result is ", string(data))

	var person2 Person
	err2 := json.Unmarshal([]byte(string(data)), &person2)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println("struct deserialize result is ", person2)
}

func serializeMap() {
	map1 := make(map[string]interface{})
	map1["name"] = "oliver"
	map1["age"] = 12
	map1["phone"] = "13222222222"
	data, err := json.Marshal(map1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("json serialize map result is ", string(data))
}

func serializeSlice() {
	slice := make([]map[string]interface{}, 2)
	slice[0] = make(map[string]interface{})
	slice[0]["name"] = "oliver"
	slice[0]["age"] = 123
	slice[0]["phone"] = "13224567213"

	slice[1] = make(map[string]interface{})
	slice[1]["name"] = "bob"
	slice[1]["age"] = 34
	slice[1]["phone"] = "14223456789"

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("json serialize slice result is ", string(data))
}
