package main

import (
	"html/template"
	"net/http"
)

func main() {
	//New Multiplexer
	mux := http.NewServeMux()
	//file server
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//handler
	mux.HandleFunc("/", index)
	mux.HandleFunc("/other", other)
	mux.HandleFunc("/struct", struc)

	//Starts the server
	http.ListenAndServe(":8080", mux)
}

//Template 1
var templ = template.Must(template.ParseFiles("index.html"))

func index(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}

//Template 2
var templ2 = template.Must(template.ParseFiles("other.html"))

func other(w http.ResponseWriter, r *http.Request) {
	templ2.Execute(w, nil)
}

//Struct
type person struct {
	Name string
	ID   string
}

//Template 3
var templ3 = template.Must(template.ParseFiles("struc.html"))

func struc(w http.ResponseWriter, r *http.Request) {
	person := person{"Robel", "ATR/4874/10"}
	templ3.Execute(w, person)
}
