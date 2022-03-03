package main

import (
	"fmt"
	"sort"
)

func main() {
	// first way to declare a slice
	var sl = []int{1, 2, 3, 4, 5}
	fmt.Println(sl)
    // adding values to the slice
	sl = append(sl, 6)

	fmt.Println(sl)

	sk := append(sl[:4]) // slice the array and append to new slice
	fmt.Println(sk)
    
	marks := []int{100,49,78,33,90,160}

	fmt.Println(marks)

	// using make method to create a slice
	vk := make([]int, 4)
	// assign values to the slice
	vk[0]  = 10
	vk[1]  = 20
	vk[2]  = 30

	fmt.Println(vk) // prints [10 20 30 0] index 3 is not assigned so it prints 0

    // sort the slice
	sort.Ints(marks)
	fmt.Println(marks)

	to_remove := 3
	// remove the value from the slice
	
	fmt.Println(marks)
	
	marks = append(marks[:to_remove], marks[to_remove+1:]...)
	fmt.Println(marks) // prints [100 49 78 33 160]



}

