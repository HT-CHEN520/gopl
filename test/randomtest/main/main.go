package main

import "github.com/HT-CHEN520/gopl/test/randomtest/utils"

//我们为了生成一个随机数,还需要个rand设置一个种子
//time.Now().unix():返回一个从1970:01:01的0时0分0秒到现在的秒数
//rand.seed(time. Now(). unix())

//如何随机的生成1-100整数
//n := rand . intn(100)+ 1   [0 100)

//fmt.Println(n)
//随机生成1-100的一个数，直到生成了99这个数，看看你一共用了几次分析思路:
//编写一个无限循环的控制，然后不停的随机生成数，当生成了99时,就退出这个无限循环==》break
func main() {
	utils.Random()
}
