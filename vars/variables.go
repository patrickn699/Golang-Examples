package main
import "fmt"


// require to use var outside the method
var allowed = true

const Pi = 3.14

func main() {
	fmt.Println("Hello, World!")

	var m int = 100
	fmt.Printf("type of m is: %T", m)
    println("\n")

	var is_n bool = true
	fmt.Printf("type of is_n is: %T", is_n)
    println("\n")

	var num uint8 = 255
	fmt.Printf("type of num is: %T", num)
    println("\n")

	var f float32 = 3.14
	fmt.Printf("type of f is: %T", f)

	var inn int
	fmt.Printf("type of inn is: %T \t",inn)
    
	//short declaration
	var abc = "abc"
	fmt.Println(abc)

	//no need to declare type
	def:= "def"
	fmt.Println(def)

	fmt.Println(allowed)
	fmt.Println(Pi)


}