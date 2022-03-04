package main

import "fmt"

func main() {

	fmt.Println("hello world")

	obj := specs{"Amd",8,"Nvidia-GTX",100000}
	fmt.Println(obj)
	// set the price
	obj.getPrice(12000)
	// price did not change
	fmt.Println(obj)

	// set the price using pointer
	obj.getPrice1(13000)
	// price changed
	fmt.Println(obj)


}

type specs struct {
	cpu string
	ram int
	gpu string
	price int
}




// setter method
func (s specs) getPrice(yur int){
	// this will not change the value of price as it makes a copy of the value
	s.price = yur
	fmt.Println(s.price)
}

func (s1 *specs) getPrice1(yur int){
	// this will change the value of price as it makes a reference of the value
	s1.price = yur
	fmt.Println(s1.price)
}

