package main


import "fmt"

func main(){

	// first way to declare a loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end // jump to end
		}
		fmt.Print(i, " ")

		end:
		fmt.Println("end")
	}
    fmt.Println("\n")
	// second way to declare a loop
	for j := range [5]int{1,2,3,4,5} {
		fmt.Print(j, " ")
	}
    
	// third way to declare a loop equivalent to while loop
	k := 0
	for k < 5 {
		if k == 3 {
			fmt.Println(k)
			break

		}
		
		
		k++
	}

}