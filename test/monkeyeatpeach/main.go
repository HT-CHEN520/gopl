package main

import "fmt"

func peach(day int) int {
	if day < 1 || day > 10 {
		fmt.Println("输入的天数不合法！")
		return 0
	}
	if day == 10 {
		return 1
	} else {
		return (peach(day+1) + 1) * 2
	}
}

func main() {
	var day int
	fmt.Println("请输入要查询的天数：")
	fmt.Scanln(&day)
	fmt.Printf("第%v天的桃子数量为%v\n", day, peach(day))
}
