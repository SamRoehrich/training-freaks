package gpx

// Track gpx track data structure
type Track struct {
	id int 
	Name string `xml:"trk>name"`
	Type string `xml:"trk>type"`
	Seg  struct {
		Points []Point `xml:"trkpt"`
	} `xml:"trk>trkseg"`
}

// Point gpx point data structure
type Point struct {
	id int
	Lat 		string `xml:"lat,attr"`
	Lon 		string `xml:"lon,attr"`
	Ele        string `xml:"ele"`
	Time       string `xml:"time"`
	Extensions struct {
		id int
		Power int `xml:"power"`
		TrackPointExtension struct {
			ATemp int `xml:"atemp"`
			HR    int `xml:"hr"`
			Cad   int `xml:"cad"`
		} `xml:"TrackPointExtension"`
	} `xml:"extensions"`
}
