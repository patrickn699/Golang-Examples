package main

import "fmt"
import "sync"

var wgg sync.WaitGroup
var sl []int 

func Ken(){

	fmt.Println("Hello World")
	wgg.Add(1)
	for i := 1; i < 100; i++ {
		sl = append(sl,i)
	}
	wgg.Done()

}
