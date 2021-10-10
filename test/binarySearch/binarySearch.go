package main

import "fmt"

//方法1
/* var findVal int

func BinarySearch(arr *[8]int) {

	var leftIndex int = 0
	var rightIndex int = len(*arr) - 1
	var midIndex int = (leftIndex + rightIndex) / 2

	for leftIndex <= rightIndex {
		if findVal == (*arr)[midIndex] {
			fmt.Printf("找到该数，下标为%v", midIndex)
			break
		} else if findVal > (*arr)[midIndex] {
			leftIndex = midIndex + 1
			midIndex = (leftIndex + rightIndex) / 2
		} else  {
			rightIndex = midIndex - 1
			midIndex = (leftIndex + rightIndex) / 2
		}
	}

	if leftIndex > rightIndex {
		fmt.Println("未找到该数字！")
	}

}

func main() {
	arr := [8]int{10, 13, 25, 58, 63, 78, 89, 97}
	fmt.Println("请输入要查询的数：")
	fmt.Scanln(&findVal)
	BinarySearch(&arr)
} */

//使用递归
func BinarySearch(arr *[8]int, leftIndex int, rightIndex int, findVal int) {

	if leftIndex > rightIndex {
		fmt.Println("未找到该数字！\n")
		return
	}

	midIndex := (leftIndex + rightIndex) / 2

	if findVal == (*arr)[midIndex] {
		fmt.Printf("找到该数，下标为%v\n", midIndex)
	} else if findVal > (*arr)[midIndex] {
		BinarySearch(arr, midIndex+1, rightIndex, findVal)
	} else {
		BinarySearch(arr, leftIndex, midIndex-1, findVal)
	}

}

func main() {
	arr := [8]int{10, 13, 25, 58, 63, 78, 89, 97}
	BinarySearch(&arr, 0, len(arr)-1, 90)
}
