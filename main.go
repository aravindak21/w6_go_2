package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Let's create a struct to represent a movie in the collection
type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
}

var movies []Movie
var nextMovieID int = 1

// Create a function for a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil || movie.Title == "" || movie.Director == "" || movie.Year == 0 || movie.Genre == " " {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	movie.ID = nextMovieID
	nextMovieID++
	movies = append(movies, movie)

	w.Header().Set("Content", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

// Get All the movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, movie := range movies {
		if movie.ID == id {
			w.Header().Set("Content", "application/json")
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

// CRUD - Update the movie using ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method Invalid", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedMovie Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil || updatedMovie.Title == "" || updatedMovie.Director == "" || updatedMovie.Year == 0 || updatedMovie.Genre == "" {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	for i, movie := range movies {

		if movie.ID == id {
			movies[i].Title = updatedMovie.Title
			movies[i].Director = updatedMovie.Director
			movies[i].Year = updatedMovie.Year
			movies[i].Genre = updatedMovie.Genre
			w.Header().Set("Content", "application/json")
			json.NewEncoder(w).Encode(movies[i])
			return
		}
	}
	http.Error(w, "Movie is unavailable", http.StatusNotFound)
}

// Delete movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Unavailable", http.StatusMethodNotAllowed)
		return
	}

	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid Movie ID", http.StatusBadRequest)
		return
	}

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid path")
	}
	return strconv.Atoi(parts[2])
}

func main() {
	//handle routes for "/movies" and "/movies/{id}"
	http.HandleFunc("/movies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMovies(w, r)
		case http.MethodPost:
			createMovie(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/movies/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getMovie(w, r)
		case http.MethodPut:
			updateMovie(w, r)
		case http.MethodDelete:
			deleteMovie(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//start server
	fmt.Println("Movie API server running successfully on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
