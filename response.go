package main

type GetActorHandlerResponse struct {
	Success bool    `json: "success"`
	Actors  []Actor `json: "actors"`
}

type AddMovieHandlerResponse struct {
	Success bool `json: "success"`
}

type AddActorsHandlerResponse struct {
	Success bool `json: "success"`
}

type DeleteMovieHandlerResponse struct {
	Success bool `json: "success"`
}
