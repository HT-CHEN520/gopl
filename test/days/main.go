//编写程序，根据输入的年份和月份，求出该月的天数（1-12）

package main

import "fmt"

func main() {
	var year int
	var month int
	fmt.Println("请输入要查询的年份：")
	fmt.Scanln(&year)
	fmt.Println("请输入要查询的月份：")
	fmt.Scanln(&month)

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Printf("%v年%v月有31天", year, month)
	case 4, 6, 9, 11:
		fmt.Printf("%v年%v月有30天", year, month)
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			fmt.Printf("%v年%v月是闰年，有29天", year, month)
		} else {
			fmt.Printf("%v年%v月是平年，有28天", year, month)
		}

	}
}
