package main

import (
	"log"
	"net/http"
)

func HttpServer() {
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
