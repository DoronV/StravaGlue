package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func combineFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("uploadPage.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		uploadFiles(r)
	} else {
		panic("No such method allowed")
	}
}

func Hello(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/combine", combineFiles)
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
