package main

import "fmt"

/*
func main() {
	var month int
	var age int

	fmt.Println("请输入月份：")
	fmt.Scanln(&month)

	if month >= 4 && month <= 10 {
		fmt.Println("请输入年龄：")
		fmt.Scanln(&age)
		if age >= 18 && age <= 60 {
			fmt.Println("票价为60元")
		} else if age < 18 {
			fmt.Println("票价为30元")
		} else if age > 60 {
			fmt.Println("票价为20")
		}
	} else {
		fmt.Println("请输入年龄：")
		fmt.Scanln(&age)
		if age >= 18 && age <= 60 {
			fmt.Println("票价为40")
		} else {
			fmt.Println("票价为20")
		}
	}
}
*/

//改进
func main() {
	var month int
	var age int

	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age > 60 {
			fmt.Printf("%v月份年龄为%v岁的票价为20元\n", month, age)
		} else if age >= 18 {
			fmt.Printf("%v月份年龄为%v岁的票价为60元\n", month, age)
		} else {
			fmt.Printf("%v月份年龄为%v岁的票价为30元\n", month, age)
		}
	} else {
		if age >= 18 && age <= 60 {
			fmt.Printf("%v月份年龄为%v岁的票价为40元\n", month, age)
		} else {
			fmt.Printf("%v月份年龄为%v岁的票价为20元\n", month, age)
		}
	}
}
