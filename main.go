package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type book struct {
	Name   string
	IsRent bool
	Id     string
}

var books = []book{}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9012", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	method := r.Method
	var encoder = json.NewEncoder(w)
	if strings.Index(r.URL.Path, "/books") == 0 {
		if r.URL.Path == "/books" && method == "GET" {
			encoder.Encode(books)
		}
		if r.URL.Path == "/books" && method == "POST" {
			postHandle(encoder, w, r)
		}
		if method == "PUT" {
			putHandle(encoder, w, r)
		}
	}
}
