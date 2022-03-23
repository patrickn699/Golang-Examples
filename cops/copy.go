package main

import ("fmt"
    "io"
	"os"
)

func main() {
	//fil,er := os.Open("imgg")
	im, _ := os.OpenFile("qkiklabs.png", os.O_RDWR|os.O_CREATE, 0666)
	fmt.Println(im)
	y, e := os.Create("immg/qk.png")
	if e != nil {
		fmt.Println(e)
	}

	str := "hello world"
	fmt.Println(str[:8])

	
		
	//k, _ :=os.Create("immg/qkiklabss.png")
	
	by, er :=io.Copy(y, im)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(by)



}