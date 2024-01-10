package main

import "fmt"

func main() {

	fmt.Println("functions")
	// another function cannot be written in function
	fmt.Print(sum(3, 5))
	fmt.Print(funres(3, 5))
	fmt.Println(proadder(2, 3, 4, 56, 3))
}

func sum(a int, b int) int {
	// datatype is written after variable eg. a int , b string etc
	return a + b
} //normal function

var funres = func(a int, b int) int {
	// datatype is written after variable eg. a int , b string etc
	return a + b
} //anothertype

func proadder(value ...int) int {
	totalsum := 0
	for i := range value {
		totalsum += value[i]
	}
	return totalsum
} //profunctions
