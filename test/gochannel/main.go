package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		rand.Seed(time.Now().Unix())
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()

	go func() {
		defer wg.Done()
		for v := range out {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}

/* func main() {

	// 声明一个等待组
	var wg sync.WaitGroup

	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	// 遍历这些地址
	for _, url := range urls {

		// 每一个任务开始时, 将等待组增加1
		wg.Add(1)

		// 开启一个并发
		go func(url string) {

			// 使用defer, 表示函数完成时将等待组值减1
			defer wg.Done()

			// 使用http访问提供的地址
			resp, err := http.Get(url)

			fmt.Println(resp)
			// 访问完成后, 打印地址和可能发生的错误
			fmt.Println(url, err)

			// 通过参数传递url地址
		}(url)
	}

	// 等待所有的任务完成
	wg.Wait()

	fmt.Println("over")
}
*/
