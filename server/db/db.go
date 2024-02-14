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

// loadEnvVariable loads and returns the value of the specified environment variable.
//
// It takes a key of type string as a parameter and returns a string.
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

// NewDatabase initializes a new Database.
//
// No parameters.
// Returns a pointer to Database and an error.
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

// Close closes the database connection.
//
// It does not take any parameters.
// It does not return any values.
func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// GetDB returns the SQL database.
//
// No parameters.
// Returns *sql.DB.
func (d *Database) GetDB() *sql.DB {
	return d.db
}
