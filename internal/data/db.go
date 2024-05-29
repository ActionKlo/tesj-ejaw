package data

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

func InitDB() {
	dsn := "postgres://admin:password@100.104.232.63:5432/ejaw?sslmode=disable"

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
