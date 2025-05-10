package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/movies", GetMovies)
	http.HandleFunc("/actors", GetActors)

	http.ListenAndServe(":8080", nil)

}
