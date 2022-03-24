package main

import "fmt"


type square struct {
	side int
}

type rect struct {
	width, height int
}

type area interface {
	area()
}

func main() {
	fmt.Println("Hello, World!")
	recta := rect{width: 10, height: 5}
	squa := square{side: 5}
	recta.area()
	squa.area()

	vi := []area{rect{20,25}, square{10}}
	for v := range vi {
		vi[v].area()
	}
}


func (s square) area(){
	fmt.Println("area of square is ",s.side * s.side)
}

func (r rect) area(){
	fmt.Println("area of rectangle is ",r.width * r.height)
}