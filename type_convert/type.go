package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inn := bufio.NewReader(os.Stdin)
	fmt.Println("type anythig: ")

	input, _ := inn.ReadString('\n')
	fmt.Println("you have entered", input)
	inp := strings.TrimSpace(input)
	//v :=  "78"
	
	k, err  := strconv.Atoi(inp)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("k is", k)
}