package db

import (
	"database/sql"
	"fmt"

	// Import postgresql library
	_ "github.com/lib/pq"
)

var (
	// DB postgresql database
	DB *sql.DB
)

// Connect to postgresql database
func Connect(host, port, sslmode, dbname, user, password string) error {
	var err error

	// connect
	connStr := fmt.Sprintf("host=%s port=%s sslmode=%s dbname=%s user=%s password=%s",
		host, port, sslmode, dbname, user, password)
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	return DB.Ping()
}
