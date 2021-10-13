package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              //指针
	slice  []int             //切片
	map1   map[string]string //map
}

func main() {
	//定义结构体变量
	//方式1
	var p1 Person

	//方式
	//p2 :=Person{}
	//或者直接定义中赋值(用的较多！！)
	//p2 := Person{"tom", 20,...}

	//方法3和方法4返回的是结构体指针

	//方式3-&
	//因为p3是一个指针,因此标准的给字段赋值方式为(*p3). Name = "smith”
	//(*p3). Name = "smith”也可以这样写 p3.Name = "smith”
	//原因:go的设计者为了程序员使用方便，底层会对 p3.Name = "smith”进行处理
	//会给3加上取值运算（*p3).Name = "smith
	//var p3 *Person = new(Person)
	//(*p3).Name = "smith"
	//p3.Name = "john"
	//fmt.Println(*p3)

	//方法4-{}
	//var p4 *Person = &Person{}
	//var p4 *Person = &Person{"mary", 20} //或直接赋值
	//fmt.Println(*p4)

	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}

	//结构体中使用切片需要先make
	p1.slice = make([]int, 5)
	p1.slice[0] = 25

	//结构体中使用map也需要先make
	p1.map1 = make(map[string]string)
	p1.map1["key1"] = "tom"

	fmt.Println(p1)
}
