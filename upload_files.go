package main

import (
	"fmt"
	"net/http"
)

func uploadFiles(r *http.Request) {
	r.ParseMultipartForm(0 << 10)
	file, handler, err := r.FormFile("gpxFile")
	if err != nil {
		fmt.Println("Something went wrong", err)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded File: ", handler.Filename)
}
