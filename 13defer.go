package main

import "fmt"

func main() {
	//defer basically means line will be executed at last
	defer fmt.Println("hello")
	fmt.Println("world")
	// world hello
	//if more than one defer statement then lifo will be followed

	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")

	// world hello one two three
}
