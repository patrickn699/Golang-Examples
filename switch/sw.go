package main

import "fmt"

func main(){

	item := []string{"apple", "banana", "orange","grapes","mango"}
	fmt.Println(item)

	
	for i := 0; i < len(item); i++ {
		
	
	switch item[i] {
		case "apple":
			fmt.Println("apple")

		case "banana":
			fmt.Println("banana")
		
		case "orange":
			fmt.Println("orange")
		
		case "grapes":
			fmt.Println("grapes")
		
		case "mango":
			fmt.Println("mango")
		
		default:
			fmt.Println("not found")
	}

}

}