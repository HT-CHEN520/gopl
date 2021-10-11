package main

import "fmt"

//使用map[string]map[string]strtng 的map类型
//key:表示用户名，是唯一的，不可以重复
//如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息，(包括昵称nickname和密码pwd） 。
//编写一个函数modifyUser(users map[string]map[string]sting,name string)完成上述功能

func modifyUser(users map[string]map[string]string, name string) {
	//判断users中是否有name
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = name
	}
}

func main() {
	users := make(map[string]map[string]string)

	//加入smith的个人信息
	users["smith"] = make(map[string]string)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "U"

	modifyUser(users, "Tom")
	modifyUser(users, "mary")

	//测试已存在用户
	modifyUser(users, "smith")

	fmt.Println(users)
}
