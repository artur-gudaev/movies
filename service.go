package main

import (
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

func GetMovie(movies []Movie) []Movie {

	sort.Slice(movies, func(i, j int) bool {
		return movies[i].Year > movies[j].Year
	})
	return movies
}

func AddMovie(movie Movie) bool {

	for index, _ := range movies {

		if movies[index].Name == movie.Name {
			return false
		}

	}

	movies = append(movies, movie)
	return true
}

func DeleteMovie(movie DeleteMovieHandlerRequest) bool {

	for number, _ := range movies {

		if movie.Name == movies[number].Name {

			movies = append(movies[:number], movies[number+1:]...)
			return true
		}
	}
	return false
}

func GetActor(movie GetActorHandlerRequest) (actor GetActorHandlerResponse) {

	for number, _ := range movies {

		if movie.Name == movies[number].Name {
			actor.Success = true
			for index, _ := range movies[number].Actors {
				actor.Actors = append(actor.Actors, movies[number].Actors[index])
			}
			break
		}
	}
	return actor
}

func AddActor(actors AddActorsHandlerRequest) bool {

	for number, _ := range movies {

		if actors.Name == movies[number].Name {

			movies[number].Actors = append(movies[number].Actors, actors.Actors...)
			return true
		}
	}

	return false
}
