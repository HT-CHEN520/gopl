package main

import "fmt"

/* 题目描述：
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。 */

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	tempx := x
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		res = res*10 + pop
	}
	//fmt.Println(res)
	return res == tempx
}

/*
思路：
如果是负数则一定不是回文数，直接返回 false
如果是正数，则将其倒序数值计算出来，然后比较和原数值是否相等
如果是回文数则相等返回 true，如果不是则不相等 false */

func main() {
	var num int
	fmt.Println("请输入要验证的数字：")
	fmt.Scanln(&num)
	result := isPalindrome(num)
	fmt.Println(result)
}

//方法2 转成字符串比较 效率更高
/* func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 &&x!=0 {
		return false
	} else {
		str := strconv.Itoa(x)
		i := 0
		for j := len(str) - 1; i < j; j-- {
			if str[i] != str[j] {
				return false
			}
			i++
		}
	}
	return true
} */
