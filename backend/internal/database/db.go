package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func ConnPostGresDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:1@localhost:5432/Todo?sslmode=disable")
		if err != nil {
			log.Fatal("Database connection failed:", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatal("Database ping failed:", err)
		}

		log.Println("Connected to Database!")
	})

	return db
}

func GetDB() *sql.DB {
	if db == nil {
		return ConnPostGresDB()
	}
	return db
}
