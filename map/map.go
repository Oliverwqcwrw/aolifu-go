package main

import "fmt"

func main() {

	map1 := make(map[string]string)
	map1["str1"] = "hello1"
	map1["str2"] = "hello2"
	delete(map1, "str1")

	map1 = make(map[string]string)
	fmt.Println(map1)

	create2()

	create3()

	createStudentInfo()

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
