package main

import (
	"encoding/json"
	"net/http"
)

func postHandle(encoder *json.Encoder, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
