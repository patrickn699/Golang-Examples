package main

import ("fmt"
    "io"
	"os"
)

func main() {
	//fil,er := os.Open("imgg")
	im, _ := os.OpenFile("qkiklabs.png", os.O_RDWR|os.O_CREATE, 0666)
	
		
	k, _ :=os.Create("immg/qkiklabss.png")
	
	by, er :=io.Copy(k, im)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(by)



}