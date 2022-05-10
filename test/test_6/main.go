package main

import "fmt"

func t1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func t2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

func main() {
	t1()
	t2()
}
