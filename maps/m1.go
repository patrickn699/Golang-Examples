package main

import "fmt"

func main() {
    
	// map is a collection of key-value pairs like a dictionary in other languages
	lang := make(map[int]string)

	// adding values to the map
	lang[1] = "Go"
	lang[2] = "Python"
	lang[3] = "Java"
	lang[4] = "C"
	lang[5] = "C++"
	// output
	fmt.Println(lang)

	delete(lang, 1) // delete the value from the map
	fmt.Println(lang)
	fmt.Println("\n")

    // iterate over the map
	// if value matches the key, then print the value
	sk := "C" 
	for k, v :=  range lang {
		if v == sk {
		fmt.Printf("the key is %d and the value is %s", k, v)
		}
	}


	



}
