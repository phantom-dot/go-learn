package main

import "fmt"

func main() {
	//3 variations of for loops

	days := []string{"mon", "tue", "thu", "fri", "sat", "sun"}

	fmt.Println(days)

	//print usig loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

}
