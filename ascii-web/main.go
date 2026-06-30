package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Create a file server handler pointing to your local directory
	fileServer := http.FileServer(http.Dir("./static"))

	// 2. Route all paths starting with "/static/" to that file server
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// FIX: Removed the empty spaces at the very beginning of the strings
	http.HandleFunc("GET /{$}", HomeHandler)
	http.HandleFunc("POST /ascii-art", ASCIIHandler)

	fmt.Println("Server starts at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
