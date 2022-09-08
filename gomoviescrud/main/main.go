package main

import (
	"fmt"
	"log"
	"movie/movie"

	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	movie.Init()
	r.HandleFunc("/movies", movie.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", movie.GetMovie).Methods("GET")
	r.HandleFunc("/movies", movie.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", movie.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", movie.DeleteMovie).Methods("DELETE")

	fmt.Printf("start the server at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
