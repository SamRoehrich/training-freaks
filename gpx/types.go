package gpx

import (
	"encoding/xml"
	"time"
)

// GPX represents a GPX file.
type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Version string   `xml:"version,attr"`
	Creator string   `xml:"creator,attr"`
	Metadata   Metadata   `xml:"metadata"`
	Waypoints  []Waypoint `xml:"wpt"`
	Tracks     []Track    `xml:"trk"`
}

// Metadata represents the metadata section of a GPX file.
type Metadata struct {
	Name        string    `xml:"name"`
	Description string    `xml:"desc"`
	Author      string    `xml:"author"`
	Time        time.Time `xml:"time"`
}

// Waypoint represents a single waypoint in a GPX file.
type Waypoint struct {
	Latitude  float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
	Elevation float64 `xml:"ele"`
	Name      string  `xml:"name"`
}

// Track represents a track in a GPX file.
type Track struct {
	Name      string     `xml:"name"`
	Number    int        `xml:"number"`
	Segments  []Segment  `xml:"trkseg"`
}

// Segment represents a segment in a GPX file.
type Segment struct {
	Points []TrackPoint `xml:"trkpt"`
}

// TrackPoint represents a single point in a GPX file's track.
type TrackPoint struct {
	Latitude  float64   `xml:"lat,attr"`
	Longitude float64   `xml:"lon,attr"`
	Elevation float64   `xml:"ele"`
	Time      time.Time `xml:"time"`
}
