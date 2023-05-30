package gpx

import "fmt"

func UploadActivity(t []byte) bool {
	track := ParseFile(t)

	fmt.Println("hi", track)


	return true
}