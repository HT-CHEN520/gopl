package main

import "fmt"

//7-14行形成一个闭包
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36) //36>>'$'
		fmt.Println("str=", str)
		return n
	}
}

func main() {

	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))

}
