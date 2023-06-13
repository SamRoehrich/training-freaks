package db

import (
	"database/sql"
	"os"
	"path/filepath"
)

// Seed creates the initial database schema
func Seed(d *sql.DB) error {

	// load seed file
	wd, err := os.Getwd()

	path := filepath.Join(wd, "db", "seed.sql")

	content, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	_, err = d.Exec(string(content))

	if err != nil {
		panic(err)
	}

	return nil
}