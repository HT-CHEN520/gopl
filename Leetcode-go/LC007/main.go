package main

import (
	"fmt"
	"math"
)

/* 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−2^31,  2^31 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号） */

func main() {
	var num int
	fmt.Println("请输入数字：")
	fmt.Scanln(&num)
	y := reverse(num)
	fmt.Printf("%v\n", y)
	//fmt.Println(math.MaxInt32)
}
func reverse(x int) int {
	ret := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		ret = ret*10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}
	return ret
}
