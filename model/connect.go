package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var con *sql.DB

func Connect() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL") ///"postgres://postgres@localhost:5432/test?sslmode=disable"
	log.Println("DB_URL: " + dbUrl)

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db.Ping())
	fmt.Println("Connected to the database")
	con = db
	return db
}
