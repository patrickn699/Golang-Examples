package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

)

var t *template.Template
var t2 *template.Template

func main(){

	mux := http.NewServeMux() // create a new ServeMux
	// render the html template
	t = template.Must(template.ParseFiles("templates/index.gohtml"))
	t2 = template.Must(template.ParseFiles("templates/details.gohtml"))
	// get the static file like css, js, images
	fs := http.FileServer(http.Dir("./static"))
	// handle the static file
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// route to the home page
	mux.HandleFunc("/", home)
	mux.HandleFunc("/res", res)	
	// start the server
	fmt.Println("server is running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", mux))

	

}



func home(w http.ResponseWriter, r *http.Request){
	//fmt.Println("home")
	//fmt.Fprintf(w, " You are at home page")
	t.Execute(w, nil)
	
	
}

func res(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "You are at res page")

	nm := r.FormValue("login")
	pass := r.FormValue("password")
	comp := nm +" "+ pass
	fmt.Println(comp)

	com := struct {
		First string
		Last string
	}{
		First: nm,
		Last: pass,
	}


	t2.ExecuteTemplate(w, "details.gohtml", com)
}