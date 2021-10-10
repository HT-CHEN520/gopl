package main

import "fmt"

func main() {

	var scores [3][5]float64
	//循环输入成绩
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%v班第%v个学生的成绩:", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	totalSum := 0.0 //用于累计所有班级的总分
	//遍历二维数组
	for i := 0; i < len(scores); i++ {
		sum := 0.0 //用于累计每个班的总分
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totalSum += sum
		fmt.Printf("第%v班的总分为%v，平均分为%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级的总分为%v，所有班级的平均分为%v\n", totalSum, totalSum/15)
}
