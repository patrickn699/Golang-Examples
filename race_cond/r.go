package main

import ("fmt"
		"sync")


func main() {
	fmt.Println("Go routines race conditions")

	wg := &sync.WaitGroup{} // create a wait group
	mute := &sync.Mutex{} // create a mutex
	var buff = []int{} // create a slice

	wg.Add(3) // add 3 routines

	go func(w *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("one")
		m.Lock() // lock the buffer
		buff = append(buff,1) // append to buffer
		m.Unlock() // unlock the buffer
		wg.Done() // notify main() that this go routine is done


	}(wg, mute)

	go func(w *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("two") 
		m.Lock() // lock the buffer
		buff = append(buff,2) // append to buffer
		m.Unlock() // unlock the buffer
		wg.Done() // notify main() that this go routine is done

	}(wg, mute) 

	go func(w *sync.WaitGroup, m *sync.Mutex){
		fmt.Println("three")
		m.Lock() // lock the buffer
		buff = append(buff,3) // append to buffer
		m.Unlock() // unlock the buffer
		wg.Done() // notify main() that this go routine is done

	}(wg, mute)


	
	wg.Wait() // wait for all go routines to finish dont close main()

	fmt.Println(buff) // print the buffer



}