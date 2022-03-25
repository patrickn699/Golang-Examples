package main

import ("fmt"
"sync")


var sg sync.WaitGroup

var inc int = 0
var lok sync.Mutex // create mutex object

func main() {
	fmt.Println("demo of mutex")
	sg.Add(1000) // add 1000 routines
	for i:=0; i< 1000; i++ {
		go compute()	
	}
	
	sg.Wait() // wait until all the routines finish

	fmt.Println("done updating:", inc)
}


func compute(){
	
	lok.Lock() // lock the current routine
	inc++ // inc the value
	lok.Unlock() // unlock the current routine
	sg.Done() // notify that the routine is finished modified the value
	
}