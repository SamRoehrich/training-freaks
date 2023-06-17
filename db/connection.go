package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
var	connStr = "postgres://default:WxU8nqpC1uVa@ep-cool-firefly-492956.us-west-2.postgres.vercel-storage.com:5432/verceldb"

// CreateConnection creates a db connection
func CreateConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Error connecting to database")
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error maintaining DB connection.")
		return nil, err
	}

	err = Seed(db)

	if err != nil {
		fmt.Println(err)
		panic("Failed to properly seed database")
	}
	return db, nil
}