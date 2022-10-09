package main

import (
	"fmt"
	"sort"
)

func main() {

	map1 := make(map[string]string)
	map1["str1"] = "hello1"
	map1["str2"] = "hello2"
	delete(map1, "str1")

	// delete all element
	map1 = make(map[string]string)
	fmt.Println(map1)

	create2()

	create3()

	createStudentInfo()

	queryElement()

	traverse()

	sortMap()

	studentMap()

}

func create2() {
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	map1["str1"] = "hello1"
	map1["str2"] = "hello2"
	fmt.Println(map1)
}

func create3() {
	var map1 = map[string]string{"str1": "hello1", "str2": "hello2"}
	fmt.Println(map1)
}

func createStudentInfo() {
	map1 := make(map[string]map[string]string)
	map1["1"] = make(map[string]string)
	map1["1"]["name1"] = "bob"
	map1["1"]["age1"] = "12"

	map1["2"] = make(map[string]string)
	map1["2"]["name2"] = "oliver"
	map1["2"]["age2"] = "23"

	fmt.Println(map1)
}

func queryElement() {
	map1 := make(map[string]string)
	map1["key1"] = "value1"

	val, ok := map1["key1"]
	if ok {
		fmt.Println(val)
	}
}

func traverse() {
	map1 := make(map[string]string)
	map1["key1"] = "value1"
	map1["key2"] = "value2"

	for k, v := range map1 {
		fmt.Printf("key = %v value = %v \n", k, v)
	}
}

func sortMap() {
	map1 := make(map[int]int)
	map1[12] = 12
	map1[23] = 23
	map1[2] = 2
	map1[1] = 1

	var arrSlice []int

	for k, _ := range map1 {
		arrSlice = append(arrSlice, k)
	}

	sort.Ints(arrSlice)

	for _, value := range arrSlice {
		fmt.Printf("map1[%v] = %v \n", value, map1[value])
	}
	fmt.Println()
}

type Stu struct {
	Id   int
	Name string
	Age  int
}

func studentMap() {
	stuMap := make(map[int]Stu)

	student1 := Stu{1, "oliver", 18}
	student2 := Stu{2, "bob", 23}

	stuMap[12] = student1
	stuMap[23] = student2

	for k, v := range stuMap {
		fmt.Printf("stuMap[%v] id=%v name=%v age=%v \n", k, v.Id, v.Name, v.Age)
	}
}
