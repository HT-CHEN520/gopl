package main

import "fmt"

func main() {
	/* 	var char byte
	   	fmt.Println("请输入一个字符")
	   	fmt.Scanf("%c", &char)

	   	switch char {
	   	case 'a':
	   		fmt.Println("A")
	   	case 'b':
	   		fmt.Println("B")
	   	case 'c':
	   		fmt.Println("C")
	   	case 'd':
	   		fmt.Println("D")
	   	case 'e':
	   		fmt.Println("E")
	   	default:
	   		fmt.Println("other")
	   	} */
	var str string = "hello，world北京"
	/* str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	} */
	str = "abc~ok北京"

	for index, val := range str {
		fmt.Printf("index = %d,val = %c \n", index, val)
	}

}
