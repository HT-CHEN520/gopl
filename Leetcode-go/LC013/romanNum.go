package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	ans := m[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if m[s[i]] >= m[s[i+1]] {
			ans += m[s[i]]
		} else {
			ans -= m[s[i]]
		}
	}
	return ans
}

func main() {
	var num string
	fmt.Println("请输入罗马数字：")
	fmt.Scanln(&num)
	y := romanToInt(num)
	fmt.Printf("%v\n", y)
	//fmt.Println(math.MaxInt32)
}
