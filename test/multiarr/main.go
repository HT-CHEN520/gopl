package main

import "fmt"

/* 用二维数组输出图形
0 0 0 0 0 0
0 0 1 0 0 0
0 0 0 0 0 0
*/

func main() {
	var arr [4][6]int

	//初始化数组
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			//只能用print
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	//fmt.Println(arr)
}
