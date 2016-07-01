package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Name   string
	IsRent bool
	Id     string
}

var books = []Book{}

func main() {
	http.HandleFunc("/books", handler)
	log.Fatal(http.ListenAndServe(":9012", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	method := r.Method
	var encoder = json.NewEncoder(w)

	if method == "GET" {
		encoder.Encode(books)
	}
	if method == "POST" {
		postHandle(encoder, w, r)
	}
	if method == "PUT" {
		encoder.Encode(books)
	}
}
