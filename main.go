package main

import (
	"fmt"
)

type Account struct {
	userName string
	passWord string
	Email    string
	Phone    string
	address  inForMation
}
type inForMation struct {
	address string
}

func main() {
	user1 := Account{

		userName: "yeulangtham1",
		passWord: "0123",
		Email:    " Hunganh@gmail.com",
		Phone:    "023666234",
		address: inForMation{
			address: "so 23, duong 3/2",
		},
	}
	user1Point := &user1 // tro den dia chi user1
	user1Point.changToPassWord("Thanh0123")
	user1.in()

}
func (user Account) in() {
	fmt.Printf("%+v\n", user)
}
func (n *Account) changToPassWord(passWord string) {
	n.passWord = passWord
}
