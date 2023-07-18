package db

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"samroehrich/training-freaks/gpx"
	"sync"
)

// InsertGpxData runs the sql queries to insert the GPX struct into the database
func InsertGpxData(gpx *gpx.GPX, db *sql.DB) (int64, error) {
	var gpxID int64
	res, err := db.Exec("INSERT INTO gpx (version, creator) VALUES (?, ?)",
		gpx.Version, gpx.Creator)

	if err != nil {
		log.Fatal("Error inserting into GPX table", err)
		return gpxID, err
	}

	gpxID, err = res.LastInsertId()

	if err != nil {
		log.Fatal("Error getting last inserted id on GPX table")
		return gpxID, err
	}

	_, err = db.Exec("INSERT INTO metadata (gpx_id, name, description, author, time) VALUES (?, ?, ?, ?, ?)",
		gpxID, gpx.Metadata.Name, gpx.Metadata.Description, gpx.Metadata.Author, gpx.Metadata.Time)
	if err != nil {
		log.Fatal("Error inserting metadata", err)
		return gpxID, err
	}

	for _, waypoint := range gpx.Waypoints {
		_, err = db.Exec("INSERT INTO waypoints (gpx_id, latitude, longitude, elevation, name) VALUES (?, ?, ?, ?, ?)",
			gpxID, waypoint.Latitude, waypoint.Longitude, waypoint.Elevation, waypoint.Name)
		
		if err != nil {
			log.Fatal("Error inserting waypoint", err)
			return gpxID, err
		}
	}

	for _, track := range gpx.Tracks {
		var trackID, segmentID int64
		res, err := db.Exec("INSERT INTO tracks (gpx_id, name, number) VALUES (?, ?, ?)",
			gpxID, track.Name, track.Number)
		if err != nil {
			log.Fatal("Error inserting track", err)
			return gpxID, err
		}

		trackID, err = res.LastInsertId()

		if err != nil {
			log.Fatal("error getting trackId", err)
			return gpxID, err
		}

		res, err = db.Exec("INSERT INTO segments (track_id) VALUES (?)", trackID)
		if err != nil {
			log.Fatal("Error inserting into segments", err)
			return gpxID, err
		}


		for _, segment := range track.Segments {
			for _, point := range segment.Points {
				fmt.Println("Points", len(segment.Points))
				_, err = db.Exec("INSERT INTO track_points (segment_id, latitude, longitude, elevation, time) VALUES (?, ?, ?, ?, ?)",
					segmentID, point.Latitude, point.Longitude, point.Elevation, point.Time)
				if err != nil {
					return gpxID, err
				}
			}
		}

		// fmt.Println("inserting points")
		// err = insertPoints(&track, &segmentID, db, 1000, 20)

	}

	fmt.Println("Insert completed")
	return gpxID, nil
}

func insertPoints(t *gpx.Track, segID *int64, db *sql.DB, batchSize, maxConcurrency int) error {
	stmt, err := db.Prepare("INSERT INTO track_points (segment_id, latitude, longitude, elevation, time) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	errChan := make(chan error)
	pointChan := make(chan []gpx.TrackPoint)
	
	var wg sync.WaitGroup
	wg.Add(maxConcurrency)
	for i := 0; i < maxConcurrency; i++ {
		// start go routines
		go func() {
			defer wg.Done()
			// loop over points arrs
			for points := range pointChan {
				tx, err := db.Begin()

				if err != nil {
					errChan <- err
					continue
				}

				// loop over each point in the array and add it to the batch
				for _, point := range points {
					fmt.Println("adding point", point.Elevation)
					_, err := tx.Stmt(stmt).Exec(*segID, point.Latitude, point.Longitude, point.Elevation, point.Time)
					if err != nil {
						errChan <- err
						break
					}
				}

				err = tx.Commit()
				if err != nil {
					errChan <- err
				}
			}
		}()
	}

	numPoints := len(t.Segments[0].Points)
	numBatches := int(math.Ceil(float64(numPoints) / float64(batchSize)))

	for i := 0; i < numBatches; i++ {
		startIndex := i * batchSize
		endIndex := (i + 1) * batchSize
		if endIndex > numPoints {
			endIndex = numPoints
		}

		pointsBatch := t.Segments[0].Points[startIndex:endIndex]
		pointChan <- pointsBatch
	}

	close(pointChan)
	
	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		fmt.Println("we have an error", err)
		if err != nil {
			return err
		}
	}

	return nil
}
