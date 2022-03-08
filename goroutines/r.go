package main

import ("fmt"
        "sync")

var vg sync.WaitGroup // create a buffered channel
var mute sync.Mutex // create a mutex to lock the buffer
var buff []int


func main() {
	
	fmt.Println("Go routines")
	go calculate_sum(10, 20)
	go calculate_sub(10, 20)
	vg.Add(2) // Add 2 go routines to wait group



	vg.Wait() // Wait for all go routines to finish dont close main()

	
}


func calculate_sum(a int, b int) {
	defer vg.Done() // notify main() that this go routine is done
	for i:=0; i<6; i++ {
		mute.Lock() // lock the buffer
		buff = append(buff, a + b)
		mute.Unlock() // unlock the buffer
	fmt.Println(a + b)
	}
}

func calculate_sub(a int, b int) {
	defer vg.Done() // notify main() that this go routine is done
	for i:=0; i<6; i++ {
		mute.Lock() // lock the buffer
		buff = append(buff, a + b)
		mute.Unlock() // unlock the buffer
	fmt.Println(a - b)
}
}