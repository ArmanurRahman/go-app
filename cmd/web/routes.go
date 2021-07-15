package main

import (
	"net/http"

	"github.com/ArmanurRahman/go-app/pkg/config"
	"github.com/ArmanurRahman/go-app/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
