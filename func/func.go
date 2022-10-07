package main

import "fmt"

func main() {

	var num int = 10
	updateValue(&num)
	fmt.Println("updateValue after is ", num)

	a := sum
	fmt.Printf("a type is %T \n", a)
	fmt.Println("a value is ", a(10, 20))

	b := paramFun(sum, 10, 20)
	fmt.Println("b value is ", b)

	customType()

	c := 10
	d := 20
	swap(&c, &d)
	fmt.Printf("c value is %v  d value is %v", c, d)

}

func updateValue(num *int) {
	*num = *num + 10
	fmt.Println("updateValue  is ", *num)
}

func sum(a int, b int) int {
	return a + b
}

func paramFun(paramFun1 func(int, int) int, num1 int, num2 int) int {
	return paramFun1(num1, num2)
}

func customType() {
	type myInt int
	var a myInt = 10
	fmt.Println("a value is ", a)
	fmt.Printf("a value type is %T \n", a)
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}
