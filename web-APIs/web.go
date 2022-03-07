package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const urll = "http://localhost:8080"

func main() {
	res, err := http.Get(urll) // get the url
	check_error(err)
	fmt.Printf("the response is %T\n:",res)

	defer res.Body.Close() // close the connection

	// read the response
	body, err := ioutil.ReadAll(res.Body)
	check_error(err)
	fmt.Println(string(body)) // convert the byte to string

    fmt.Print("\n")
	fmt.Print("\n")



}


func check_error(err error) {
	if err != nil {
		panic(err)
	}
}