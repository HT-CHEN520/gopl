package binarysearch

import "fmt"

var findVal int

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
		} else {
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
}
