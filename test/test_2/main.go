package main

import "fmt"

func main() {
	var a float64 = 11.0
	var b float64 = 18.5
	if a > 10.0 && b < 20.0 {
		fmt.Println("和=", a+b)
	}
}
