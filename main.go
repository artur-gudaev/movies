package main

import (
	//"fmt"
	"encoding/json"
	"net/http"
)

var movies = map[string][]string{}

type AddMovieHandlerRequest struct {
	Movie  string   `json: movie`
	Actors []string `json: actors`
}

type AddMovieHandlerResponse struct {
	Success bool `json: success`
}

type DeleteMovieHandlerRequest struct {
	Movie string `json: movie`
}

type DeleteMovieHandlerResponse struct {
	Success bool `json: success`
}

type PutMovieHandlerRequest struct {
	Movie  string   `json: movie`
	Actors []string `json: actors`
}

type PutMovieHandlerResponse struct {
	Success bool `json: success`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

func AddMoviesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req AddMovieHandlerRequest
	var resp AddMovieHandlerResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, ok := movies[req.Movie]; ok {
		resp.Success = true
	} else {
		resp.Success = false
		movies[req.Movie] = req.Actors
	}

	json.NewEncoder(w).Encode(resp)
}

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req DeleteMovieHandlerRequest
	var resp DeleteMovieHandlerResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, ok := movies[req.Movie]; ok {
		delete(movies, req.Movie)
		resp.Success = true
	} else {
		resp.Success = false
	}

	json.NewEncoder(w).Encode(resp)
}

func PutMovieHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req PutMovieHandlerRequest
	var resp PutMovieHandlerResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if _, ok := movies[req.Movie]; ok {
		movies[req.Movie] = append(movies[req.Movie], req.Actors...)
		resp.Success = true
	} else {
		resp.Success = false
	}

	json.NewEncoder(w).Encode(resp)

}
func main() {

	movies["Terminator"] = []string{"Scvarcnereg"}

	http.HandleFunc("/get", GetMovies)
	http.HandleFunc("/post", AddMoviesHandler)
	http.HandleFunc("/delete", DeleteMovieHandler)
	http.HandleFunc("/put", PutMovieHandler)

	http.ListenAndServe(":8080", nil)

}
