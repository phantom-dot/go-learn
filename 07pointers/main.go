package main

import "fmt"

func main() {
	fmt.Println("pointers	")

	var a *int
	fmt.Println("value of pntr is ", a)

	mynum := 23
	var ptr = &mynum
	fmt.Println("value of ptr is ", ptr)
	fmt.Println("value of ptr is ", *ptr)
}
