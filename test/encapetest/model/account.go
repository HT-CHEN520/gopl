package model

import "fmt"

type account struct {
	AccountNo string
	pwd       string
	balance   float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对")
	}
	return &account{
		AccountNo: accountNo,
	}
}

func (a *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对")
	} else {
		a.pwd = pwd
	}

}

func (a *account) GetPwd() string {
	return a.pwd
}

func (a *account) SetBalance(balance float64) {
	if balance < 20 {
		fmt.Println("余额账目不对")
	} else {
		a.balance = balance
	}
}

func (a *account) GetBalance() float64 {
	return a.balance
}
