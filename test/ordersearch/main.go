package main

import "fmt"

func main() {
	name := [5]string{"中", "国", "欢", "迎", "您"}
	var searchName string
	fmt.Println("请输入要搜索的字")
	fmt.Scanln(&searchName)

	//方法1
	/* for i := 0; i < len(name)-1; i++ {
		if searchName == name[i] {
			fmt.Printf("找到%v,下标为%v\n", searchName, i)
			break
		} else if i == (len(name) - 1) {
			fmt.Printf("没有找到%v \n", searchName)
		}
	} */

	//方法2
	index := -1
	for i := 0; i < len(name)-1; i++ {
		if searchName == name[i] {
			index = i
			break
		}
	}

	if index != -1 {
		fmt.Printf("找到%v,下标为%v\n", searchName, index)
	} else {
		fmt.Printf("没有找到%v \n", searchName)
	}

}
