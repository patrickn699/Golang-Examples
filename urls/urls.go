package main

import (
	"fmt"
	"net/url"
)

const urls = "http://localhost:8080/home?name=john&age=20"

func main() {

	result, _ := url.Parse(urls)
	fmt.Println(result) // get the url
	fmt.Println("the scheme is: ",result.Scheme) // get the scheme http or https
	fmt.Println("your host is:",result.Host) // get the host
	fmt.Println("the path is:",result.Path) // get the path or route
	fmt.Println("the query is:",result.RawQuery) // get the query string in url
	fmt.Println("the port is: ",result.Port()) // get the port number

	query := result.Query() // get the query string in url

	fmt.Println("the name is:",query["name"]) // get the name from query string
	fmt.Println("the age is:", query["age"])  // get the age from query string


	build_url := &url.URL{
		Scheme: "http",
		Host: "localhost:8080",
		Path: "/home",
		RawQuery: "name=john&age=20",

	}

	fmt.Println("the constructed url is:",build_url)

	str_url := build_url.String()
	fmt.Println("the string url is:",str_url)


}

