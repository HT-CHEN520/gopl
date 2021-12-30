package main

import (
	"fmt"

	"github.com/HT-CHEN520/gopl/test/encapetest/model"
)

func main() {
	a := model.NewAccount("adgjags")
	a.SetPwd("sgsag~")
	a.SetBalance(30)
	fmt.Println(a)
	fmt.Println(a.AccountNo, a.GetPwd(), a.GetBalance())
}
