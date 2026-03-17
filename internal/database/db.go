package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {

	connStr := "host=localhost port=5432 user=admin password=admin dbname=inventory sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database connection failed")
	}

	DB = db
	log.Println("Database connected")
}
