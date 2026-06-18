package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

type PageData struct {
	Result string
	Text   string
	Banner string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error parsing file", http.StatusInternalServerError)
	}
}

func ASCIIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if text == "" || banner == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	realbanner := banner + ".txt"
	fmt.Println(text, banner)
	result, err := Generator(text, realbanner)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	Data := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}
	err = tpl.Execute(w, Data)
	if err != nil {
		http.Error(w, "Error parsing file", http.StatusInternalServerError)
	}
}

func SwitchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-switch" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	//fmt.Println(text, banner)
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if text == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if banner == "" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	realbanner := banner + ".txt"
	fmt.Println(text, banner)
	result, err := Generator(text, realbanner)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	Data := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}
	err = tpl.Execute(w, Data)
	if err != nil {
		http.Error(w, "Error parsing file", http.StatusInternalServerError)
	}
}
func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/ascii-art", ASCIIHandler)
	http.HandleFunc("/ascii-art-switch", SwitchHandler)
	http.ListenAndServe(":8080", nil)
}
