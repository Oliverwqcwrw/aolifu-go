package main

import "fmt"

func main() {

	key := ""
	exit := false

	for {
		fmt.Println("记账软件")
		fmt.Println("1: 收支明细")
		fmt.Println("2: 登记收入")
		fmt.Println("3: 登记支出")
		fmt.Println("4: 退出软件")
		fmt.Println("请选择1-4")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("收支明细记录")
		case "2":
		case "3":
			fmt.Println("登记支出")
		case "4":
			exit = true
		default:
			fmt.Println("请输入正确的选项")
		}

		if exit {
			break
		}

	}
}
