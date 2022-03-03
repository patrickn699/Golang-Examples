package main

import "fmt"

func main() {

	var ar [4]string

	ar[0] = "Hello"
	ar[1] = "World"
	ar[3] = "Go"

	fmt.Println(ar)
	fmt.Println(len(ar))

	// second way to declare an array
	ar2 := [4]string{"Hello", "World", "Go", "Lang"}
	fmt.Println(ar2[:4]) // slice the array
}