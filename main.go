package main

import (
	"log"
	"os"

	"samroehrich/training-freaks/gpx"
)

func main() {
		data, err := os.ReadFile("sample-gpx/power.gpx")

		if err != nil {
			log.Fatal(err)
		}
		
		gpx.ParseFile(data)
}