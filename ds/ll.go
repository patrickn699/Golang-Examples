package main

import "fmt"

type node struct {
	value int
	next *node
}

func createNode(value int) *node {
	return &node{value: value}
}

func (n *node) add(value int) {
	n.next = createNode(value)
}

func (n *node) print() {
	fmt.Println(n.value)
}

func (n *node) printAll() {
	for n != nil {
		n.print()
		fmt.Println(n)
		n = n.next
	}
}

func (n *node) remove(value int) {
	if n.value == value {
		n.next = n.next.next
	} else {
		n.next.remove(value)
	}
}


func main() {
	fmt.Println("create a linked list")
	head := createNode(2)
	h2 := createNode(3)
	head.add(9)
	h2.add(4)
	head.printAll()
	//h2.printAll()
	fmt.Println("remove 3")
	head.remove(2)
	head.printAll()
}

