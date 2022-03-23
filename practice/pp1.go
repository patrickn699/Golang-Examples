package main

import "fmt"
import "sync"

func main() {
	k := make(map[int]string)
    v := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	
	for i := 1; i < 27; i++ {
		k[i] = v[i-1]
	}

	fmt.Println(k)
	keep()
	//pp2.Ken()
}


func keep(){
	var wg sync.WaitGroup
	for i :=1; i < 27; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

}

