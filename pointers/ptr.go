package main

import "fmt"


func main() {
	fmt.Println("Hello, playground")
	
	var k *int // pointer to an int
	
	fmt.Println(k) // prints nil

	v := 50 // int variable
	l := &v // pointer to v
	fmt.Println(*l) // prints 50

	*l = *l + 5 // v = v + 5
	fmt.Println(v) // prints 55

	temp := 10
	fmt.Println("before updating:", temp)
	update(temp)
	fmt.Println("after updating:",temp)


}

func update(s int){
	s = s + 1
	fmt.Println("inside update:",s)
}