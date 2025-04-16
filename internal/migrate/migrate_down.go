package main

import (
	"github.com/fote15/go-url-shortener/internal/database"
	"github.com/joho/godotenv"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	err := godotenv.Load()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed: ", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // path to migrations folder
		"postgres",          // DB dialect
		driver,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("✅ Database down migrated successfully")
}
