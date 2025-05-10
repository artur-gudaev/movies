package main

import (
	"encoding/json"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:

		w.Header().Set("Content-Type", "application/json")

		movies = GetMovie(movies)

		json.NewEncoder(w).Encode(movies)

	case http.MethodPost:

		w.Header().Set("Content-Type", "application/json")

		var req Movie
		var resp AddMovieHandlerResponse

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		resp.Success = AddMovie(req)

		json.NewEncoder(w).Encode(resp)

	case http.MethodDelete:

		w.Header().Set("Content-Type", "application/json")

		var req DeleteMovieHandlerRequest
		var resp DeleteMovieHandlerResponse

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		resp.Success = DeleteMovie(req)

		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func GetActors(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		var req GetActorHandlerRequest
		var resp GetActorHandlerResponse

		req.Name = r.URL.Query().Get("name")

		resp = GetActor(req)

		json.NewEncoder(w).Encode(resp)

	case http.MethodPost:

		w.Header().Set("Content-Type", "application/json")

		var req AddActorsHandlerRequest
		var resp AddActorsHandlerResponse

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		resp.Success = AddActor(req)

		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
