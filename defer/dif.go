package main

import "fmt"	


// defer will be executed after the function returns
// in reverse order
func main(){
	tot:=0
	for i:= range [5]int{1,2,3,4,5} {
		tot += i
		defer fmt.Println(tot)
	}
}