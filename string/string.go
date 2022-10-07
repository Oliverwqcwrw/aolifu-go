package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "hello"
	fmt.Println("str length is ", len(str))

	// 按字节计算长度
	str = "hello中"
	fmt.Println("str length is ", len(str))

	arr := []rune(str)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("char = %c \n", arr[i])
	}

	str2Int()

	i2Str()

	strToByte()

	byte2Str()

	scaleConvert()

	isContain()

	strCount()

	equalsIgnore()

	firstIndexOf()

	lastIndexOf()

	replaceStr()

	splitStr()

	lowerCase()

	upperCase()

	trimSpace()

	trimSpecifiedStr()

	startsWithOrEndWith()

}

func str2Int() {
	str := "123"
	num, error := strconv.Atoi(str)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println("num value is ", num)
	}
}

func i2Str() {
	num := 12345
	str := strconv.Itoa(num)
	fmt.Println("str value is ", str)
}

func strToByte() {
	str := "hello world"
	var bytes = []byte(str)
	fmt.Println("bytes is", bytes)
}

func byte2Str() {
	str := string([]byte{97, 98, 99})
	fmt.Println("str value is ", str)
}

func scaleConvert() {
	num := 123
	str1 := strconv.FormatInt(int64(num), 2)
	fmt.Printf("num 的二进制 is %v \n", str1)
	str2 := strconv.FormatInt(int64(num), 16)
	fmt.Printf("num 的十六进制 is %v \n", str2)
}

func isContain() {
	flag := strings.Contains("abc", "ab")
	fmt.Println("flag value is ", flag)
}

func strCount() {
	count := strings.Count("abcdee", "e")
	fmt.Println("count value is ", count)
}

func equalsIgnore() {
	fmt.Println(strings.EqualFold("abc", "Abc"))
	fmt.Println("abc" == "Abc")
}

func firstIndexOf() {
	fmt.Println(strings.Index("abcdefg", "def"))
}

func lastIndexOf() {
	fmt.Println(strings.LastIndex("abcdabc", "ab"))
}

func replaceStr() {
	fmt.Println(strings.Replace("go go hello world", "go", "golang", 1))
}

func splitStr() {
	str := "hello,world,hello"
	arr := strings.Split(str, ",")
	fmt.Println("arr value is ", arr)
}

func lowerCase() {
	fmt.Println(strings.ToLower("HELLO"))
}

func upperCase() {
	fmt.Println(strings.ToUpper("hello"))
}

func trimSpace() {
	fmt.Println(strings.TrimSpace(" hello world "))
}

func trimSpecifiedStr() {
	fmt.Println(strings.Trim("!hello world!", "!"))

	fmt.Println(strings.TrimLeft("!hello world !", "!"))

	fmt.Println(strings.TrimRight("!hello world!", "!"))
}

func startsWithOrEndWith() {
	fmt.Println(strings.HasPrefix("!hello world !", "!"))

	fmt.Println(strings.HasSuffix("!hello world!", "!"))
}
