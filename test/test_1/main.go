package main

import "fmt"

func main() {
	var time float64
	var gender string
	fmt.Println("请输入参赛选手的时间：")
	fmt.Scanln(&time)
	if time < 8 {
		fmt.Println("请输入参赛选手的性别：")
		fmt.Scanln(&gender)
		if gender == "男" {
			fmt.Println("进入决赛的男子组")
		} else {
			fmt.Println("进入决赛的女子组")
		}
	} else {
		fmt.Println("淘汰")
	}
}
