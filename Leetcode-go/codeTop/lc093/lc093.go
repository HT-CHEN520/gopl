package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "25525511135"
	res := restoreIpAddresses(s)
	fmt.Print(res)
}

func restoreIpAddresses(s string) []string {
	ans := []string{}
	var dfs func(subRes []string, start int)
	dfs = func(subRes []string, start int) {
		if len(subRes) == 4 && start == len(s) {
			ans = append(ans, strings.Join(subRes, "."))
			return
		}
		if len(subRes) == 4 && start < len(s) {
			return
		}

		for length := 1; length <= 3; length++ {
			if start+length-1 >= len(s) {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			if n, _ := strconv.Atoi(str); n > 255 {
				return
			}
			subRes = append(subRes, str)
			dfs(subRes, start+length)
			subRes = subRes[:len(subRes)-1]
		}
	}
	dfs([]string{}, 0)
	return ans
}
