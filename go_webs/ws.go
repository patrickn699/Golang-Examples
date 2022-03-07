package main

import ("fmt"
        "log"
       "net/http")

func main() {

	fmt.Println("Running the server on port 8081")
	fmt.Println("Server is running and the url is http://localhost:8081")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from docker container")
	} )

	log.Fatal(http.ListenAndServe(":8081", nil))


	

}