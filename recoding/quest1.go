package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

type Page struct {
	Text string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/echo" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	text := r.FormValue("text")
	Echo := Page{
		Text: text,
	}
	if err := tpl.Execute(w, Echo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//fmt.Fprintln(w, text)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/echo", EchoHandler)
	fmt.Println("Starting a server at  port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
