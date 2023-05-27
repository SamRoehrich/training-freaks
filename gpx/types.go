package gpx

type Track struct {
	Name string `xml:"trk>name"`
	Type string `xml:"trk>type"`
	Seg  struct {
		Points []Point `xml:"trkpt"`
	} `xml:"trk>trkseg"`
}

type Point struct {
	Lat 		string `xml:"lat,attr"`
	Lon 		string `xml:"lon,attr"`
	Ele        string `xml:"ele"`
	Time       string `xml:"time"`
	Extensions struct {
		Power int `xml:"power"`
		TrackPointExtension struct {
			ATemp int `xml:"atemp"`
			HR    int `xml:"hr"`
			Cad   int `xml:"cad"`
		} `xml:"TrackPointExtension"`
	} `xml:"extensions"`
}
