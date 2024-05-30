package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("postgres://%s:%s@db:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	var err error
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed connect to database with err:", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database")
}
