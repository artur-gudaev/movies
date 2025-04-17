package main

import (
	"encoding/json"
	"net/http"
	"sort"
)

type Actor struct {
	NameRole map[string]string `json: "namerole"`
	Age      int               `json: "age"`
}
type Movie struct {
	Name   string  `json: "name"`
	Year   int     `json: "year"`
	Actors []Actor `json: "actors"`
}

var movies = []Movie{}

type GetActorHandlerRequest struct {
	Name string `json: "name"`
}

type GetActorHandlerResponse struct {
	Success bool    `json: "success"`
	Actors  []Actor `json: "actors"`
}

type AddMovieHandlerResponse struct {
	Success bool `json: "success"`
}

type AddActorsHandlerRequest struct {
	Name   string  `json: "name"`
	Actors []Actor `json: "actors"`
}

type AddActorsHandlerResponse struct {
	Success bool `json: "success"`
}

type DeleteMovieHandlerRequest struct {
	Name string `json: "name"`
}

type DeleteMovieHandlerResponse struct {
	Success bool `json: "success"`
}

type PutMovieHandlerRequest struct {
	Movie  string   `json: "movie"`
	Actors []string `json: "actors"`
}

type PutMovieHandlerResponse struct {
	Success bool `json: "success"`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].Year > movies[j].Year
	})

	json.NewEncoder(w).Encode(movies)
}

func GetActors(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req GetActorHandlerRequest
	var resp GetActorHandlerResponse
	found := false
	var number int

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for number, _ = range movies {

		if req.Name == movies[number].Name {

			found = true
			break

		}
	}

	if found {

		resp.Success = true
		for index, _ := range movies[number].Actors {
			resp.Actors = append(resp.Actors, movies[number].Actors[index])
		}
	} else {
		resp.Success = false
	}

	json.NewEncoder(w).Encode(resp)

}
func AddMoviesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req Movie
	var resp AddMovieHandlerResponse
	found := false

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for index, _ := range movies {

		if movies[index].Name == req.Name {
			found = true
			break
		}

	}

	if found {

		resp.Success = false

	} else {

		movies = append(movies, req)
		resp.Success = true

	}

	json.NewEncoder(w).Encode(resp)
}

func AddActorsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req AddActorsHandlerRequest
	var resp AddActorsHandlerResponse
	found := false
	var number int

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for number, _ = range movies {

		if req.Name == movies[number].Name {

			found = true
			break
		}
	}

	if found {

		resp.Success = true
		movies[number].Actors = append(movies[number].Actors, req.Actors...)

	} else {

		resp.Success = false

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
	found := false
	var number int

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for number, _ = range movies {

		if req.Name == movies[number].Name {

			found = true
			break

		}

	}

	if found {

		movies = append(movies[:number], movies[number+1:]...)
		resp.Success = true

	} else {

		resp.Success = false

	}

	json.NewEncoder(w).Encode(resp)
}

func main() {

	http.HandleFunc("/get", GetMovies)
	http.HandleFunc("/getactors", GetActors)
	http.HandleFunc("/post", AddMoviesHandler)
	http.HandleFunc("/postactors", AddActorsHandler)
	http.HandleFunc("/delete", DeleteMovieHandler)

	http.ListenAndServe(":8080", nil)

}
