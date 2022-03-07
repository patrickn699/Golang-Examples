package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main(){
	// this is get request

	//GetRequest()

	// this is post request

	PostRequest()
}

func GetRequest() {
	const url = "http://localhost:8080/get"

	res, err := http.Get(url) // get the url
	check_error(err)

	defer res.Body.Close()

	fmt.Println("the status code is: ",res.Status) // get the status code
	fmt.Println("the status is: ",res.StatusCode) // get the status code
	fmt.Println("the content length is: ",res.ContentLength) // get the content length


	contens, err := ioutil.ReadAll(res.Body) // read the response
	check_error(err)

	//fmt.Print(string(contens)) // convert the byte to string shortcut way

	// second way of reading the response and converting it to string 
	// using built in strings library
	var str_content strings.Builder
    fmt.Print("\n")
	str_content.Write(contens) // convert the byte to string
	//fmt.Println("the content length is: ",cnt) // get the content length
	fmt.Println("the content is: ",str_content.String()) // convert the byte to string
	





}

func PostRequest() {
	const url = "http://localhost:8080/post"

	requestBody := strings.NewReader(`{"name":"john","age":30}`)

	res, err := http.Post(url, "application/json", requestBody) // post the url
	check_error(err)

	defer res.Body.Close()

	//fmt.Println("the status code is: ",res.Status) // get the status code
	//fmt.Println("the status is: ",res.StatusCode) // get the status code
	//fmt.Println("the content length is: ",res.ContentLength) // get the content length

	content, _ := ioutil.ReadAll(res.Body) // read the response

	fmt.Println("",string(content)) // convert the byte to string
}

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}