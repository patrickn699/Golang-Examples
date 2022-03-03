package main


import "fmt"

type My struct {
	name string
	age int
	dept string
}

func main() {
	user := My{"Patrick", 25, "IT"}
	fmt.Println(user)

	user1 := My{"Nandu", 28, "IT"}

	fmt.Println(user1.name)
}

