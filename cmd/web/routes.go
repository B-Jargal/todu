package main

import (
	"net/http"

	"github.com/B-Jargal/todu.git/pkg/application"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *application.Application) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/pub", func(r chi.Router) {
		r.Get("user", cacheUserID(app))
	})

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./ui/dist")))
	filesDir := http.Dir("/")
	FileServer(router, "/", filesDir)
	router.Mount("/", mux)
	return router
}
