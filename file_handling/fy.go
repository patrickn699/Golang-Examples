package main

import (
	"fmt"
	"io"
	"os"
)


func main(){
	fmt.Print("file handling in go\n")
    // sample string
	sample := "hello.txt, world this will be sent to the file in local directory"

	// creating a file
    fil, err := os.Create("./hello.txt")
	check_error(err)
	// closing the file using defer
	defer fil.Close()
	// writing to the file
	io.WriteString(fil, sample)
	fmt.Print("file created\n")

	// reading the file
	fmt.Print("reading the file\n")
	read_file("./hello.txt")
	
}

// function to read the file
func read_file(file string){

	fil, err := os.Open(file)
	check_error(err)
	defer fil.Close()
    
	// reading the file
	content, err := os.ReadFile(file)
	check_error(err)
	//printing the content
	fmt.Println(string(content))

	//seek the file
	file_seek(file, 5, 0)
	
}

func file_seek(file string, offset int64, whence int) {
	fil, err := os.Open(file)
	check_error(err)
	seeked, _ := fil.Seek(offset, whence)
	fmt.Println("seeking the file to the required position", seeked)
	//fmt.Println(seeked)

}

// function to check the error
func check_error(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
