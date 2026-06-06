package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

type Greet struct {
	First  string
	Second string
}

func main() {
	http.HandleFunc("/", Hello)
	http.Handle()
	http.Handler(func(w, http.ResponseWriter, r *http.Request))
	http.ServeMux()
	http.ListenAndServe(":8080", nil)
	// const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	// fmt.Println(html.EscapeString(s))
}

// var tpl template.
func Hello(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello There")
	tpl, _ = template.ParseFiles("templates/index.html")
	message := "Welcome To Divine's Corner"
	message2 := "What Could We Help You With?"
	response := Greet{message, message2}
	tpl.Execute(w, response)
}
