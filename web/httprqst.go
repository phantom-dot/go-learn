package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://www.google.com"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//fmt.Printf(response)  as it is a pointer cant print
	// we use reader
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
