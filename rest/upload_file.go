package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadFile writes a file to the system
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

	// TODO - remove the temp file
	tFile.Write(fileBytes)
	fmt.Println("TEmp File creation successful")


	err = uploadToS3(&fileBytes)

	if err != nil {
		fmt.Println("from the caller, you failed")
	}
	
	w.WriteHeader(200)
}

// TODO - fix this
// TODO - don't hardcode this
func uploadToS3(b *[]byte) error {
	creds := credentials.NewStaticCredentials("", "asdf", "asdf")
	provider := aws.CredentialsProvider(creds)
	if err != nil {
		fmt.Println("fuck me")
		return err
	}

	config := &aws.Config{
		Region: "us-east-1",
		Credentials: provider,
	}
	// sess, err := session.NewSession(config)
	
	
	if err != nil {
		return err
	}

	svc := s3.New(sess)

	input := &s3.PutObjectInput{
		Bucket: aws.String("training-freaks"),
		Key: aws.String("first.gpx"),
		Body: b,
	}

	_, err = svc.PutObject(input)

	if err != nil {
		fmt.Println("error uploading to S3")
		return err
	}

	fmt.Println("yeahhhhhh")
	return nil


}