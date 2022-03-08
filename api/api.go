package main

import "fmt"

func main() {

	fmt.Println("Hello World")

}

type workout_plan struct {

	Day string `json:"day"`
	Exercise string `json:"exercise"`
	Sets int `json:"sets"`
	Reps int `json:"reps"`
	trainer *trainer



}

type trainer struct {

	Name string `json:"name"`
	Contact int `json:"contact"`


}

