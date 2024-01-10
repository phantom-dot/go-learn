package main

import "fmt"

func main() {

	//methods are functions in structs
	sahil := user{
		Name:     "sahil",
		age:      20,
		verified: false,
		email:    "asdfas",
	}

	sahil.getstatus()

}

type user struct {
	Name     string
	age      int
	verified bool
	email    string
	// var with first letter upper case is exportable
	// else not

}

func (u user) getstatus() {
	fmt.Println("is user active", u.verified)
}
