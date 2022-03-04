package main

import "fmt"

func main() {
	fmt.Println("this is the main function")

	sec()

	fmt.Println(sum(10, 20))

	fmt.Println(sum_def(10, 20, 30, 40, 50))
    a := 10
	b := 20

	fmt.Println(sum_ptr(&a, &b))
}

func sec(){
	fmt.Println("this is the sec function with no return and args")
}

// method to sum two numbers as args
func sum(a int, b int) int{
	fmt.Println("this is the sum function with return and args")
	return a + b

}

// this is the sum function with default args or mutliple args
func sum_def(vals ...int) int{

	fmt.Println("this is the sum function with multiple args")
	total := 0
	for _, val := range vals {
		total += val
	}
	return total

}

// this is the sum function with pointer args
func sum_ptr (a , b *int) int{
	fmt.Println("this is the sum function with pointers")
	return *a + *b

}

