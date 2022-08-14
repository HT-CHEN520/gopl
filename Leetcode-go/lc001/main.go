package main

import "fmt"

func Searchadd(nums *[4]int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {

	var nums = [4]int{2, 7, 11, 15}
	var target int = 17
	var res []int = Searchadd(&nums, target)

	fmt.Printf("%v", res)
}
