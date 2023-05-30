package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)
var	connStr = "postgres://default:I4fjAJke6pBo@ep-odd-wildflower-707311.us-west-2.postgres.vercel-storage.com:5432/verceldb"

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

	return db, nil
}