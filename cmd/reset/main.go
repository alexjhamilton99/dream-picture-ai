package main

import (
	"database/sql"
	"dream-picture-ai/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func createDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	return db.CreateDatabase(dbname, user, password, host)
}

func main() {
	db, err := createDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tables := []string{
		"schema_migrations",
		"accounts",
		"images",
	}

	for _, table := range tables {
		query := fmt.Sprintf("DROP TABLE IF EXISTS %s CASCADE", table)
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}
