package main

import "fmt"

func main() {
	fmt.Println("maps")
	lang := make(map[string]int)
	//syntax => make(map[key datatype]value datatype)
	lang["hindi"] = 1
	lang["eng"] = 2
	lang["french"] = 3

	fmt.Println(lang)          //  whole list
	fmt.Println(lang["hindi"]) //1
	println("sahil")

	delete(lang, "eng")
	fmt.Println(lang)

	//for loop
	for key, value := range lang {
		fmt.Printf("for key %v,value is %v\n", key, value)

	}
}
