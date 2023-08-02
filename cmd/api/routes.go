package main

import (
	"github.com/anras5/go-with-docker/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)

	return mux
}
