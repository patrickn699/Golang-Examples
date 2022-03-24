package main

import "fmt"


func main(){
    
	k := make(chan string, 2)

	get_msg(k)

	select {
		case msg := <-k:
			fmt.Println("recieved msg:",msg)
		
		case msg2 := <-k:
			fmt.Println("recieved msg:",msg2)	

		default:
			fmt.Println("no msg recieved")
	}
  
} 


func get_msg(k chan string){
	fmt.Println("sending a msg to main")
	k <- "hello msg"
	k <- "hello msg2"
}