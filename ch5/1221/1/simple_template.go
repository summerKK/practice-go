package main

import (
	"net/http"
	"html/template"
)

type Page struct {
	Title, Content string
}

func main() {

	http.HandleFunc("/", displayPage)
	http.ListenAndServe(":8080", nil)
}

func displayPage(resp http.ResponseWriter, req *http.Request) {
	p := &Page{Title: "template", Content: "Hello World"}
	t := template.Must(template.ParseFiles("ch5/1221/1/template.html"))
	t.Execute(resp, p)
}
