package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *AppConfig) http.Handler {

	mux := chi.NewRouter() //declaring my new router

	//going to use one of chi package's middleware ,the recoverer(that gracefully handles error?)
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(WritetoConsole) //calling our write to function in middleware file

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)

	return mux

}
