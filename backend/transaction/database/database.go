package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func open() (*sql.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	pwd := os.Getenv("POSTGRES_PASSWORD")
	name := os.Getenv("POSTGRES_CONTAINER")
	db := os.Getenv("POSTGRES_DB")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pwd, name, db)

	return sql.Open("postgres", connStr)
}

func Open() *sql.DB {
	for {
		db, err := open()
		if err != nil {
			log.Printf("failed to create database connection: %s", err)
			goto failure
		}
		err = db.Ping()
		if err != nil {
			log.Printf("failed to ping database: %s", err)
			goto failure
		}

		log.Println("connected to database")

		return db

	failure:
		time.Sleep(2 * time.Second)
	}
}
