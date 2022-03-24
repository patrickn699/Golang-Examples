package main


import "fmt"


func main() {
	le := 2
	fmt.Println("Hello, playground")
	fmt.Println("I am player 1")
	fmt.Println("hitted the ball")
	hit := make(chan int,le)
	for i := 0; i < le; i++ {
		hit <- i
	}
	go playtennis(le, hit)
	fmt.Println("player 1 is done")
    op1 := <-hit
	op2 := <-hit
	if op1 == op2 {
		fmt.Println("I am player 2")
		fmt.Println("hitted the ball sent 4")
	}else{
		fmt.Println("I am player 2")
		fmt.Println("hitted the ball sent 5",op1,op2)
	}
}

func playtennis(le int, hit chan int) {
	v :=0
	for i := 0; i < le; i++ {
		fmt.Println("player 2 is done")
		v = <-hit
		fmt.Println("recieved",v)
		hit <- v + 1
	}
	close(hit)


	fmt.Println("playing tennis")
	fmt.Println("hitting back to player 1")
	//rec := <-hit
	
}