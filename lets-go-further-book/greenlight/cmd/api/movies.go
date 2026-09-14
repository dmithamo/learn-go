package main

import (
	"fmt"
	"net/http"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "POSTing to %s\n", r.URL)
}

func (appp *application) getMoviesHandler(w http.ResponseWriter, r * http.Request){}

func (app *application) getMovieByIdHandler(w http.ResponseWriter, r *http.Request){
	id := readParamByName("id", r)

	fmt.Fprintf(w, "show the details of movie with id %s\n", id)
}
