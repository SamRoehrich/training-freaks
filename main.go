package main

import (
	"fmt"

	"samroehrich/training-freaks/db"

	_ "github.com/lib/pq"
)

func oldmain() {

		fmt.Println("main service started...")
		
		d, err := db.CreateConnection()

		if err != nil {
			fmt.Println("DB connection error")
			return
		}

		fmt.Println("Connected to database.")

		row, err := d.Query("SELECT * FROM activity WHERE name='first'")

		if err != nil {
			fmt.Println("Error Select * from activity")
		}

		var name string
		var details string
		var category string

		for row.Next() {
			row.Scan(&name, &details, &category)
		}

		fmt.Println(name, "name")
		fmt.Println(details, "details")
	}