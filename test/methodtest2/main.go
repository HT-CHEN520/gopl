package main

import "fmt"

type MethodUtils struct {
}

func (mu MethodUtils) Print() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

/*
	1)编写结构体(MethodUtils)，编程一个方法，方法不需要参数，
	在方法中打印一个10*8 的矩形，在main方法中调用该方法。
*/
func main() {
	var mu MethodUtils
	mu.Print()
}
