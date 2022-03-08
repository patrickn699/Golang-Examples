package main

import (
	"fmt"
	"sync"
)



func main(){
	fmt.Println("This is a channel")

	ch := make(chan int, 2) // buffered channel with 2 elements
	//ch <- 10
	//fmt.Println(<-ch)
    wg := &sync.WaitGroup{} // create a wait group
    wg.Add(2) // add two to the wait group

	// chan<- means that we can only send data to the channel
	go func(ch chan<- int, wg *sync.WaitGroup){
		fmt.Println("passing the value to channel...")
		//close(ch)
		ch <- 10 // pass the value to the channel
		ch <- 20 // pass the value to the channel
		close(ch) // close the channel
		wg.Done() // notify the wait group that the goroutine is done


	}(ch,wg) // pass the channel and wait group to the goroutine


	// <-chan means that the channel is only used for reading
	func(wg *sync.WaitGroup, ch <-chan int){
		fmt.Println("getting the value from the channel...")
		// close(ch)
		val, is_channel_open := <- ch // get the value from the channel

		if is_channel_open {
			fmt.Println(val) // print the value if channel is open
		}else{
			fmt.Println("channel is closed")
		}

		//fmt.Println(<-ch) // get the value from the channel
		wg.Done() // notify the wait group that the goroutine is done
	}(wg,ch) // pass the wait group and channel to the goroutine

	wg.Wait() // wait for all go routines to finish dont close main()

}