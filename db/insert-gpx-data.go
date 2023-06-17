package db

import (
	"database/sql"
	"samroehrich/training-freaks/gpx"
)

// InsertGpxData runs the sql queries to insert the GPX struct into the database
func InsertGpxData(gpx *gpx.GPX, db *sql.DB) error {
	var gpxID int64
	err := db.QueryRow("INSERT INTO gpx (version, creator) VALUES ($1, $2) RETURNING id",
		gpx.Version, gpx.Creator).Scan(&gpxID)

	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO metadata (gpx_id, name, description, author, time) VALUES ($1, $2, $3, $4, $5)",
		gpxID, gpx.Metadata.Name, gpx.Metadata.Description, gpx.Metadata.Author, gpx.Metadata.Time)
	if err != nil {
		return err
	}

	for _, waypoint := range gpx.Waypoints {
		_, err = db.Exec("INSERT INTO waypoints (gpx_id, latitude, longitude, elevation, name) VALUES ($1, $2, $3, $4, $5)",
			gpxID, waypoint.Latitude, waypoint.Longitude, waypoint.Elevation, waypoint.Name)
		
		if err != nil {
			return err
		}
	}

	for _, track := range gpx.Tracks {
		var trackID, segmentID int64
		err := db.QueryRow("INSERT INTO tracks (gpx_id, name, number) VALUES ($1, $2, $3) RETURNING id",
			gpxID, track.Name, track.Number).Scan(&trackID)
		if err != nil {
			return err
		}

		err = db.QueryRow("INSERT INTO segments (track_id) VALUES ($1) RETURNING id", trackID).Scan(&segmentID)
		if err != nil {
			return err
		}

		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				_, err = db.Exec("INSERT INTO track_points (segment_id, latitude, longitude, elevation, time) VALUES ($1, $2, $3, $4, $5)",
					segmentID, point.Latitude, point.Longitude, point.Elevation, point.Time)
				if err != nil {
					return err
				}
			}
		}

	}

	return nil
}