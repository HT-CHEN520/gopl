package main

import "fmt"

func main() {
	//s := "babad"
	var s string
	fmt.Scan(&s)
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expendCenter(s, i, i)
		left2, right2 := expendCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expendCenter(s string, left int, right int) (int, int) {
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			break
		}
	}
	return left + 1, right - 1
}
