package router

import (
	"database/sql"
	handlers "github.com/fote15/go-url-shortener/internal/api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(db *sql.DB) http.Handler {
	urlH := handlers.NewUrlHandler(db)
	userH := handlers.NewUserHandler(db)
	r := mux.NewRouter()

	RegisterUserRoutes(r, userH)
	RegisterURLRoutes(r, urlH)

	return r
}
