package main

import (
	"fmt"

	"github.com/CodeHanHan/DataStructure-go/lichang/BinarySearch"

	cBinarySearch "github.com/HT-CHEN520/gopl/main/BinarySearch"
)

func main() {
	fmt.Printf("search: %d\n", BinarySearch.BinarySearch([]int{1, 2, 3, 4, 5}, 5))
	fmt.Printf("search: %d\n", cBinarySearch.BinarySearch([]int{1, 2, 3, 4, 5}, 3))
}
