package gpx

import (
	"encoding/xml"
	"fmt"
)

// ParseFile takes a raw GPX file and parses it out to a Track
func ParseFile(raw []byte) Track {
	var t Track
	xml.Unmarshal(raw, &t)

	fmt.Println(t.Seg.Points)

	return t
}