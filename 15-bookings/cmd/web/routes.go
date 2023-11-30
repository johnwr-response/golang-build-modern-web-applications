package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/pkg/config"
	"github.com/johnwr-response/golang-build-modern-web-applications/15-bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	if !app.InProduction {
		// This is just to make the git commit check to shut up and will be removed
	}

	return mux
}
