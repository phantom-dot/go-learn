package main

import "fmt"

func main() {
	fmt.Println("struct in go lang")
	// no inheritance in go lang and classes

	sahil := User{"sahil", "blah@balh.com", true, 69}
	fmt.Println(sahil)
	fmt.Printf("value of sahil as a user are => %+v \n", sahil)
	//%+v will give the naming and types also
	//%v will only give the values
	fmt.Printf("sahil's status is %v", sahil.status)
}

type User struct {
	Name   string
	email  string
	status bool
	Age    int
}
