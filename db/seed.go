package db

import (
	"database/sql"
)

// Seed creates the initial database schema
func Seed(d *sql.DB) error {

	createTablesStatement := []string{
		`CREATE TABLE IF NOT EXISTS gpx(
			id SERIAL PRIMARY KEY,
			version VARCHAR(255),
			creator VARCHAR(255)
			)`,
		`CREATE TABLE IF NOT EXISTS metadata(
			id SERIAL PRIMARY KEY,
			gpx_id INT,
			name VARCHAR(255),
			description VARCHAR(255),
			author VARCHAR(255),
			time TIMESTAMP,
			FOREIGN KEY (gpx_id) REFERENCES gpx (id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS waypoints(
			id SERIAL PRIMARY KEY,
			gpx_id INT,
			latitude DOUBLE PRECISION,
			longitude DOUBLE PRECISION,
			elevation DOUBLE PRECISION,
			name VARCHAR(255),
			FOREIGN KEY (gpx_id) REFERENCES gpx (id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS tracks(
			id SERIAL PRIMARY KEY,
			gpx_id INT,
			name VARCHAR(255),
			number INT,
			FOREIGN KEY (gpx_id) REFERENCES gpx (id) ON DELETE CASCADE			
		)`,
		`CREATE TABLE IF NOT EXISTS segments (
			id SERIAL PRIMARY KEY,
			track_id INT,
			FOREIGN KEY (track_id) REFERENCES tracks (id) ON DELETE CASCADE
		)`,
		`CREATE TABLE IF NOT EXISTS track_points (
			id SERIAL PRIMARY KEY,
			segment_id INT,
			latitude DOUBLE PRECISION,
			longitude DOUBLE PRECISION,
			elevation DOUBLE PRECISION,
			time TIMESTAMP,
			FOREIGN KEY (segment_id) REFERENCES segments (id) ON DELETE CASCADE
		)`,
		}

	for _, statement := range createTablesStatement {
		_, err := d.Exec(statement)

		if err != nil {
			return err
		}
	}
	return nil
}