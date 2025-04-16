// main.go
package main

import (
	"github.com/fote15/go-url-shortener/internal/api/router"
	"log"
	"net/http"
	"os"

	"github.com/fote15/go-url-shortener/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	defer db.Close()

	handler := router.SetupRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server started at port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
