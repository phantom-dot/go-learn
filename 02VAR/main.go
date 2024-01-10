package main

import "fmt"

const pi = 3.14             //public
const Logintoken = "sdfsfs" //private
func main() {

	var x string = "Hello World"
	fmt.Println(x)

	var name string = "Hello World"
	fmt.Println(name)
	fmt.Printf("var is of type %T", name)

	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	//implict type
	var d = "ehllo"
	fmt.Println(d)

	//no var
	e := 34
	fmt.Println(e)

	fmt.Print(pi)
	fmt.Print(Logintoken)
}
