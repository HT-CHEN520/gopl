package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSort(nums, 0, n-1, n-k)
}

func quickSort(nums []int, low int, high int, k int) int {
	pivotIndex := parition(nums, low, high)
	if pivotIndex == k {
		return nums[k]
	} else if pivotIndex > k {
		return quickSort(nums, low, pivotIndex-1, k)
	} else {
		return quickSort(nums, pivotIndex+1, high, k)
	}
}

func parition(nums []int, low int, high int) int {
	key := nums[low]
	for low < high {
		for low < high && nums[high] >= key {
			high--
		}
		nums[low] = nums[high]
		for low < high && nums[low] <= key {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = key
	return low
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	res := findKthLargest(nums, 2)
	fmt.Print(res)
}
