package main

import "fmt"

func main() {
	var number int
	fmt.Println("请输入一个三位数：")
	fmt.Scanln(&number)

	a := number / 100
	b := (number % 100) / 10
	c := number % 10

	if a*a*a+b*b*b+c*c*c == number {
		fmt.Println("该数为水仙花数")
	} else {
		fmt.Println("不是水仙花数")
	}
}
