package main

import (
	"encoding/json"
	"fmt"
)


func main(){

	get_json()
	Decode_json()

}


type laptop struct {
	Name string `json:"name"` // tag name to lowercase
	Cpu string  `json:"cpu"`
	Ram int     `json:"ram"`
	Os []string  `json:"os,omitempty"` // omitempty to ignore empty array
	Password string `json:"-"` // hide the password
}

func get_json(){

	lappy := []laptop {
		       {"Dell","Intel",8,[]string{"Windows","Linux"},"123"},
               {"Lenovo","Amd",6, []string{"Windows","Linux"},"456"},
			   {"Apple","M1",8, []string{"Unix"},"980"},
	}

	// convert the struct to json
	json_data, err := json.MarshalIndent(lappy,"","\t")
	check_error(err)
	fmt.Println(string(json_data))




}

func check_error(err error) {
	if err != nil {
		panic(err)
	}
}


func Decode_json(){

	my_data := []byte(`
	{
		"Name": "Apple",
		"Cpu": "M1",
		"Ram": 8,
		"Os": ["Unix"]
}`)

	var lapy laptop // create a struct

	valid := json.Valid(my_data) // check if the json is valid

	if valid  {
		fmt.Println("the json is valid: ",valid)
		json.Unmarshal(my_data,&lapy) // unmarshal the json to struct
		fmt.Printf("%#v\n",lapy) // format the struct
		
	}

    fmt.Print("\n")

	var dict map[string]interface{} // create a map
	json.Unmarshal(my_data,&dict) // unmarshal the json to map
	//fmt.Printf("%#v\n",dict) // format the map
	fmt.Println(dict) // get the value of the map




	

}


