package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// CreateConnection creates a db connection
func CreateConnection() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	dsn := os.Getenv("DSN") 
	db, err := sql.Open("mysql", dsn)

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
	fmt.Println("successfully connected to PlanetScale")
	return db, nil
}