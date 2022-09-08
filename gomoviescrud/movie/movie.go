package movie

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies map[string]Movie

func Init() {
	movies = make(map[string]Movie)
	movies["1"] = Movie{Isbn: "10000", Title: "Jaws", Director: &Director{Firstname: "Steven", Lastname: "Spielberg"}}
	movies["2"] = Movie{Isbn: "10001", Title: "Jaws 2", Director: &Director{Firstname: "Jeannot", Lastname: "Szwarc"}}
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	setContentType(w)
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	setContentType(w)
	params := mux.Vars(r)

	index := params["id"]
	if _, ok := movies[index]; !ok {
		json.NewEncoder(w).Encode("Could not delete movie with index " + index)
	} else {
		delete(movies, index)
		json.NewEncoder(w).Encode(movies)
	}
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	setContentType(w)
	params := mux.Vars(r)
	index := params["id"]

	if _, ok := movies[index]; !ok {
		json.NewEncoder(w).Encode("Could not find movie with index " + index)
	} else {
		json.NewEncoder(w).Encode(movies[index])
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies[strconv.Itoa(rand.Intn(10000))] = movie
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	setContentType(w)

	params := mux.Vars(r)
	index := params["id"]

	if _, ok := movies[index]; !ok {
		json.NewEncoder(w).Encode("Could not update movie with index " + index)
	} else {
		var movie Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movies[index] = movie
		json.NewEncoder(w).Encode(movie)
	}
}

func setContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
