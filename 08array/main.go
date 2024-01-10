package main

import "fmt"

func main() {
	var array [5]int //declartion
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println("array is ", array)
	fmt.Println("array lengtj is ", len(array))

	var array2 = [3]int{}
	fmt.Print(array2)

}
