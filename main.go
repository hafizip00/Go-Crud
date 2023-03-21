package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}


func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type" , "application/json")
	params := mux.Vars(r)
	for index  , item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, v := range movies {
		if v.ID == params["id"]{
			json.NewEncoder(w).Encode(v)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var  movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(1000000)))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, movie := range movies{
		if movie.ID == params["id"]{
			movies = append(movies[:index] , movies[index:]... )
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies , movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}


func main() {
	r := mux.NewRouter()


	movies = append(movies , Movie{ID : "1" , Isbn : "123456" , Title : "MOVIE1" , Director : &Director{FirstName : "John" , LastName : "Doe"}})
	movies = append(movies , Movie{ID : "2" , Isbn : "654321" , Title : "MOVIE2" , Director : &Director{FirstName : "John" , LastName : "Smith"}})
	

	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovie).Methods("GET")
	r.HandleFunc("/movies" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}" , deleteMovie).Methods("DELETE")

	fmt.Println("Starting... at 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}