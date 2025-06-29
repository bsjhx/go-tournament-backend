package db

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func RunMigrations() {
	log.Println("Running database migrations...")

	db, err := sql.Open("sqlite3", "tournament.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("Failed to set goose dialect: %v", err)
	}

	// Standard migrations
	if err := runMigrationsFromPath(db, "migrations"); err != nil {
		log.Fatalf("Standard migrations failed: %v", err)
	}

	// Private migrations
	privatePath := os.Getenv("PRIVATE_MIGRATIONS_PATH")
	if privatePath != "" {
		if err := runMigrationsFromPath(db, privatePath); err != nil {
			log.Fatalf("Private migrations failed: %v", err)
		}
	}

	log.Println("Migrations completed successfully")
}

func runMigrationsFromPath(db *sql.DB, path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Printf("Migration path does not exist: %s (skipping)", path)
		return nil
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	log.Printf("Applying migrations from: %s", absPath)
	return goose.Up(db, absPath)
}
