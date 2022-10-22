package main

import (
	"fmt"
	"reflect"
)

const (
	nameTest = "Hello World"
	ageTest  = 12
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (stu Student) GetSum(a int, b int) int {
	return a + b
}

func (stu Student) Print() {
	fmt.Println(stu)
}
func main() {
	a := 10
	checkType(a)

	stu := Student{
		Name: "oliver",
		Age:  12,
	}
	checkStructType(stu)

	testConst1()

	var num1 int = 10
	modifyValue(&num1)
	fmt.Println("num1 value is ", num1)

	stu2 := Student{
		Name: "bob",
		Age:  23,
	}
	getStructField(stu2)

	getStructMethod(stu2)
}

func checkType(a interface{}) {
	rType := reflect.TypeOf(a)
	fmt.Println(rType)

	rValue := reflect.ValueOf(a)
	fmt.Println(rValue)

	fmt.Printf("rType type is %T, rValue type is %T \n", rType, rValue)

	num1 := 2 + rValue.Int()
	fmt.Println(num1)

	rI := rValue.Interface()
	num2 := rI.(int)
	fmt.Println(num2)
}

func checkStructType(obj interface{}) {
	rT := reflect.TypeOf(obj)
	rValue := reflect.ValueOf(obj)
	rI := rValue.Interface()
	stu := rI.(Student)
	fmt.Println(stu.Name)

	rTK := rT.Kind()
	rVK := rValue.Kind()
	fmt.Println("type kind is ", rTK, "value kind is \n", rVK)

}

func testConst1() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}

func modifyValue(num interface{}) {
	rV := reflect.ValueOf(num)
	rV.Elem().SetInt(20)
}

func getStructField(obj interface{}) {
	rV := reflect.ValueOf(obj)
	rVK := rV.Kind()

	if rVK != reflect.Struct {
		fmt.Errorf("the type is not struct")
	}
	fieldNum := rV.NumField()
	fmt.Println("field num is ", fieldNum)

	rT := reflect.TypeOf(obj)
	for i := 0; i < fieldNum; i++ {
		fmt.Printf("field [%v] value is %v \n", i, rV.Field(i))

		tagName := rT.Field(i).Tag.Get("json")
		if tagName != "" {
			fmt.Printf("field [%v] tag is %v \n", i, tagName)
		}

	}
}

func getStructMethod(obj interface{}) {
	rV := reflect.ValueOf(obj)
	methodNum := rV.NumMethod()
	fmt.Println("obj method  num is", methodNum)

	rV.Method(1).Call(nil)

	valueSlice1 := make([]reflect.Value, 2)
	valueSlice1[0] = reflect.ValueOf(10)
	valueSlice1[1] = reflect.ValueOf(20)
	respSlice := rV.Method(0).Call(valueSlice1)
	fmt.Println(respSlice[0].Int())

}
