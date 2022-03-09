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
    
	// create a new ServeMux
	mux := http.NewServeMux() 
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

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
    
	// get the form values	
	nm := r.FormValue("login")
	pass := r.FormValue("password")
	comp := nm +" "+ pass

	fmt.Println(comp)

	// pass the values to the template through a struct
	com := struct {
		First string
		Last string
	}{
		First: nm,
		Last: pass,
	}

	//fmt.Println(com)
    
	/*
	type data struct{
		First string
		Last string
	}
    
	dat := data{nm, pass}
	*/
    // execute the template and pass the struct
	t2.ExecuteTemplate(w, "details.gohtml", com)
	}
}