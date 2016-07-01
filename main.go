package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Book struct {
	Name   string
	IsRent bool
	Id     string
}

var books = []Book{}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/books", handler)
	log.Fatal(http.ListenAndServe(":9012", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	newBook := Book{
		Name:   RandStringRunes(13),
		IsRent: false,
		Id:     RandStringRunes(8),
	}
	books = append(books, newBook)
	var encoder = json.NewEncoder(w)
	encoder.Encode(books)
}
