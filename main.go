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
		r.ParseMultipartForm(0 << 10)
		file, handler, err := r.FormFile("gpxFile")
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		defer file.Close()
		fmt.Println("Uploaded File: %+v\n", handler.Filename)
	} else {
		panic("No such method allowed")
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/combine", combineFiles)
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
