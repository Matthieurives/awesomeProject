package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	Port = ":8080"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./html/test.html")
	if err != nil {
		fmt.Println(err)
	}
	items := struct {
		Pokemon string
		Type    string
	}{
		Pokemon: "pikachu",
		Type:    "elec",
	}
	t.Execute(w, items)
}

func main() {
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(Port, nil)
}
