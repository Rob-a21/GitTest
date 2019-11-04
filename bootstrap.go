package main

import (
	"html/template"
	"net/http"
)

type first struct {
	Name string
	ID   string
}

var templ = template.Must(template.ParseFiles("index.html"))
var tem = template.Must(template.ParseFiles("other.html"))
var t = template.Must(template.ParseFiles("struc.html"))

func structServe(w http.ResponseWriter, r *http.Request) {
	user := first{"Robel", "ATR/4874/10"}
	t.Execute(w, user)
}

func index(w http.ResponseWriter, r *http.Request) {
	templ.Execute(w, nil)
}

func other(w http.ResponseWriter, r *http.Request) {
	tem.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	//File Server
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//Handler
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/other", other)
	mux.HandleFunc("/struct", structServe)
	//Listener
	http.ListenAndServe(":8000", mux)
}
