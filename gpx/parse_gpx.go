package gpx

import (
	"encoding/xml"
)

// ParseFile takes a raw GPX file and parses it out to a Track
func ParseFile(raw []byte) Track {
	var t Track
	xml.Unmarshal(raw, &t)

	return t
}