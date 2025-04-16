package router

import (
	"github.com/fote15/go-url-shortener/internal/api/handlers"
	middleware "github.com/fote15/go-url-shortener/internal/api/middleware"
	"github.com/gorilla/mux"
)

func RegisterURLRoutes(r *mux.Router, h *handlers.UrlHandler) {
	r.HandleFunc("/url/{shortKey}", h.Redirect).Methods("GET")

	api := r.PathPrefix("/urls").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)

	api.HandleFunc("/shorten", h.ShortenURL).Methods("POST")
	api.HandleFunc("/{id}", h.DeleteURL).Methods("DELETE")
	api.HandleFunc("/{id}", h.EditURL).Methods("PUT")
	api.HandleFunc("/{id}/stats", h.GetStats).Methods("GET")
	api.HandleFunc("/", h.ListUserURLs).Methods("GET")
}
