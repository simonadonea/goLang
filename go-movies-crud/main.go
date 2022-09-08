package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go-movies-crud/movie"
)

// type Movie struct {
// 	// ID       string    `json:"id"`
// 	Isbn     string    `json:"isbn"`
// 	Title    string    `json:"title"`
// 	Director *Director `json:"director"`
// }
// type Director struct {
// 	Firstname string `json:"firstname"`
// 	Lastname  string `json:"lastname"`
// }

var movies map[string]Movie

func main() {
	r := mux.NewRouter()

	movies = make(map[string]Movie)
	movies["1"] = Movie{Isbn: "10000", Title: "Jaws", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}}
	movies["2"] = Movie{Isbn: "10001", Title: "Jaws 2", Director: &Director{Firstname: "Jeannot", Lastname: "Szwarc"}}

	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("start the server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
