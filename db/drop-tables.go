package db

import "database/sql"

// DropGpxTables drops all the tables in the database
func DropGpxTables (db *sql.DB) error {
	tables := []string{
		"track_points",
		"segments",
		"waypoints",
		"metadata",
		"tracks",
		"gpx",
	}

	_, err := db.Exec("SET session_replication_role = 'replica'")
	if err != nil {
		return err
	}

	for i := len(tables) - 1; i >= 0; i-- {
		_, err := db.Exec("DROP TABLE IF EXISTS " + tables[i])
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("SET session_replication_role = 'origin")
	if err != nil {
		return err
	}
	return nil
}