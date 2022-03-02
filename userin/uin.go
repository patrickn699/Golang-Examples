package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Print("enter anythig: ")
	from_u := bufio.NewReader(os.Stdin)

	input, _ := from_u.ReadString('\n')


	println("you have entered",input)


	

}