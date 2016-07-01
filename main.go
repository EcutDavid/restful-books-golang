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
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("r.PostForm", r.PostForm)
		log.Println("r.Form", r.Form)
		if len(r.Form["name"]) == 1 {
			newBook := Book{
				Name:   string(r.Form["name"][0]),
				IsRent: false,
				Id:     RandStringRunes(8),
			}
			books = append(books, newBook)
			encoder.Encode(books)
		}
		return
	}
	if method == "PUT" {
		encoder.Encode(books)
	}
}
