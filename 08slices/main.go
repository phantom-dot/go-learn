package main

import "fmt"

func main() {
	var slice []int
	//slice doesn't have initial size declation
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(slice)

	//slice of slice
	slice1 := slice[1:3]
	fmt.Println(slice1)

	//using make function

	slice2 := make([]int, 10)
	//array of 10 element

	slice2[0] = 1
	//so make initilize ans sets array with first 10 or given size elemnts and if we use append then the size of slice or array will be changed

	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7, 8)
	//now after this the size of slice2 will be 10+8=18

	fmt.Println(slice2)

	//remove a value from slice
	var courses = []string{"eng", "hindi", "maths", "gvgv", "tyfytf", "uygvyug"}
	courses = append(courses, "sci")
	courses = append(courses[:0], courses[1:]...)
	fmt.Println(courses)

}
