package main

import (
	"log"
	"net/http"
	"text/template"
)

// Home method renders the main.html
func Home(writer http.ResponseWriter, reader *http.Request) {
	templateHtml := template.Must(template.ParseFiles("main.html"))
	templateHtml.Execute(writer, nil)
}

// main
func main() {
	log.Println("Server started on: http://localhost:8000")
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8000", nil)
}
