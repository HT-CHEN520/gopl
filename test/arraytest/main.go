package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//1)创建一个byte类型的26个元素的数组,分别放置'A'-'z'。
	//使用for循环访问所有元素并打印出来。提示:字符数据运算'A'+1 ->'B"
	var myChar [26]byte
	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i) //注意需要将i转成byte类型
	}

	for i := 0; i < 26; i++ {
		fmt.Printf("%c", myChar[i])
	}

	fmt.Println()
	//求一个数组中的最大值
	var intArr = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}

	}
	fmt.Printf("maxVal=%v,maxValIndex=%v\n", maxVal, maxValIndex)

	//求出一个数组的和和平均值
	var intArr2 = [...]int{1, -1, 9, 90, 12}
	sum := 0
	for _, val := range intArr2 {
		sum += val
	}
	fmt.Printf("sum=%v 平均值=%v\n", sum, float64(sum)/float64(len(intArr2)))

	//随机生成5个数，并将其反转打印
	var intArr3 [5]int
	len := len(intArr3)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr3[i] = rand.Intn(100)
	}
	fmt.Println("交换前=", intArr3)
	//反转打印，交换的次数是len/2
	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr3[len-1-i]
		intArr3[len-1-i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后=", intArr3)

}
