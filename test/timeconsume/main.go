package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	str := ""
	for i := 0; i < 1000000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
	//fmt.Println(str)
}

func main() {
	start := time.Now().Unix() //开始时的时间
	test()
	end := time.Now().Unix() //结束时的时间
	fmt.Printf("执行test()耗费的时间为%v秒\n", end-start)
}
