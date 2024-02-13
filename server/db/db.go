package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Database struct {
	db *sql.DB
}

func loadEnvVariable(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("key %s not set", key)
	}
	return value
}

func NewDatabase() (*Database, error) {
	databaseUrl := loadEnvVariable("DATABASE_URL")
	if databaseUrl == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &Database{db: db}, nil

}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
