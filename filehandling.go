package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("file handling")
	content := "data to write in the file"
	file, err := os.Create("./myfile.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	fmt.Println(len(content))
	readFile("./myfile.txt")
}

func readFile(filename string) {
	//data, err := ioutil.ReadFile(filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//data is in bytes
	fmt.Println("text in file is \n", string(data))
}
