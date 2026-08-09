package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initializes PostgreSQL connection pool and verifies connection
func InitDB(dataSourceName string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(5)

	log.Println("PostgreSQL connection established successfully")
	return DB, nil
}

// GetDB returns the active database instance
func GetDB() *sql.DB {
	return DB
}
