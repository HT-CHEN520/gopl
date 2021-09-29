package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Random() {
	var count int = 0

	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		fmt.Println("n=", n)
		count++
		if n == 99 {
			break //表示跳出for循环
		}
	}
	fmt.Println("生成 99 一共使用了", count)
}
