package main

import (
	"fmt"
	"sync"
)

var sg sync.WaitGroup

func main() {
	players := []string{"miguel","dimitri","robby","sam"}
	sg.Add(len(players))
	for _,i := range players{
		go addo(i)
	}
	
	
	sg.Wait()
	fmt.Println("All team kicked")
}


func addo(player string) {
	fmt.Println(player,"used high jump kick")
	sg.Done()
}