package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    
	// reading from stdin
	inn := bufio.NewReader(os.Stdin)
	fmt.Println("type anythig: ")

	input, _ := inn.ReadString('\n')
	fmt.Println("you have entered", input)
	// removing spaces from stdin for further conversion
	inp := strings.TrimSpace(input)
	// converting string to int
	k, err  := strconv.Atoi(inp)
	// checking for error
	if err != nil {
		fmt.Println("error", err)
	}
	//printing the converted int
	fmt.Println("k is", k)
}