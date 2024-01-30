package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//PerformGetRqst("http://localhost:8000/get")
	//PerforPostJSONRqst("http://localhost:8000/post")
	//PerforPostJSONRqst("http://localhost:8000/postform")
	decodejson()
}

func PerformGetRqst(url string) {
	response, err := http.Get(url)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
}

func PerforPostJSONRqst(url string) {
	rqstbody := strings.NewReader(`
	{
		"name":"sahil",
		"age":"90",
		"email":"lohransahil@gmail.com"
	}

`) // created json to send
	response, err := http.Post(url, "application/json", rqstbody)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
}
func PerformPostFORMRqst(myurl string) {
	//formdata
	data := url.Values{}
	data.Add("firstname", "sahil")
	data.Add("age", "20")
	data.Add("email", "lohransahil@gmail.com")

	//post rqst
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// read response
	bodydata, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodydata))

}

func decodejson() {
	// Create a GET request
	req, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	defer req.Body.Close()
	// Read the response body
	body, err := ioutil.ReadAll(req.Body)
	checkjsonisvalidornot := json.Valid(body)
	if checkjsonisvalidornot {
		fmt.Println("json is valid")
		fmt.Printf("%#v\n", body)
	} else {
		fmt.Println("json not valid")
	}
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
