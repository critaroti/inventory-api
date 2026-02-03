package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed connect DB:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB not reachable:", err)
	}

	DB = db
	log.Println("Database connected")
}
