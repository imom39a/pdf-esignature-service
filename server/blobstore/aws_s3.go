package blobstore

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	S3_REGION   = "us-east-1"
	S3_ENDPOINT = "http://localhost:4572"
	S3_BUCKET   = "mytestbucket"
)

func getSession() *session.Session {
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(S3_REGION),
		Endpoint:         aws.String(S3_ENDPOINT),
	})
	if err != nil {
		log.Fatal(err)
	}

	return s
}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(file multipart.File, handler *multipart.FileHeader) error {

	uploader := s3manager.NewUploader(getSession())

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(handler.Filename),
		Body:   file,
	})

	if err != nil {
		// Print the error and exit.
		fmt.Println("Unable to upload %s to %s, %v", handler.Filename, S3_BUCKET, err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", handler.Filename, S3_BUCKET)

	return err
}

func AddFileToS3_(absoluteFilePath string, fileName string) error {

	file, err := os.Open(absoluteFilePath)
	if err != nil {
		return fmt.Errorf("failed to open file %q, %v", absoluteFilePath, err)
	}
	uploader := s3manager.NewUploader(getSession())

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		// Print the error and exit.
		fmt.Println("Unable to upload %s to %s, %v", fileName, S3_BUCKET, err)
	}

	fmt.Printf("Successfully uploaded %s to %s\n", fileName, S3_BUCKET)

	return err
}

func GetFileFromS3(fileName string) (*os.File, error) {
	tempFile, err := ioutil.TempFile("", "download-*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	downloader := s3manager.NewDownloader(getSession())
	_, err = downloader.Download(tempFile, &s3.GetObjectInput{
		Bucket: aws.String(S3_BUCKET),
		Key:    aws.String(fileName),
	})
	if err != nil {
		// Do your error handling here
		return nil, err
	}
	return tempFile, nil
}
