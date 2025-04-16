package router

import (
	"github.com/fote15/go-url-shortener/internal/api/handlers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, h *handlers.UserHandler) {
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", h.Register).Methods("POST")
	auth.HandleFunc("/login", h.Login).Methods("POST")
}
