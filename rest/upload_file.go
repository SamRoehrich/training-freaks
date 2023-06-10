package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Upload file writes a file to the system
// TODO - swap implementation to upload file to external storage
func UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Initiating the file upload... ")

	r.ParseMultipartForm(10 << 20)

	file, _, err := r.FormFile("file")

	if err != nil {
		fmt.Println("Error retreiving file from form.", err)
		return
	}

	defer file.Close()

	tFile, err := ioutil.TempFile("temp_gpx", "upload_*.gpx")

	if err != nil {
		fmt.Println("error creating temp file", err)
		return 
	}

	defer tFile.Close()

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println("error converting file to a byte array", err)
		return
	}

	tFile.Write(fileBytes)
	fmt.Println("File upload successful")

	w.WriteHeader(200)
}