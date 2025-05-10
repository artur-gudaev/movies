package main

type GetActorHandlerRequest struct {
	Name string `json: "name"`
}

type AddActorsHandlerRequest struct {
	Name   string  `json: "name"`
	Actors []Actor `json: "actors"`
}

type DeleteMovieHandlerRequest struct {
	Name string `json: "name"`
}
