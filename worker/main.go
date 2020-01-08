package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go"
	"log"
	"net/url"
	"os"
	"time"
)

var db *sql.DB
func init() {
	var err error

	pgDbName := os.Getenv("POSTGRES_DB")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")

	fmt.Println(pgDbName, pgUser, pgPass)

	connStr := "host=postgresql user=" + pgUser + " password=" + pgPass + " dbname=" + pgDbName + " sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connect to your database")
}


func main () {

	// Called function download
	Error , ErrorDescription , s3urllink := saveToS3("transcoder-out" , "ru-east-1" ,"ubuntu-18.04.3-live-server-amd64.iso" ,  "/Users/xander/Downloads/ubuntu-18.04.3-live-server-amd64.iso" ,"application/x-iso9660-image" )

	fmt.Println(Error)
	fmt.Println(ErrorDescription)
	fmt.Println(s3urllink)
}

func saveToS3 (bucketName string , location string , objectName string , filePath string , contentType string) (Error int, ErrorDescription string , s3urllink string) {

	// If the IP and PORT data are incorrect, an error is generated only after the timeout expires ~ 90 сек

	s3host := os.Getenv("S3_HOST")
	s3port := os.Getenv("S3_PORT")
	s3accessKey := os.Getenv("MINIO_ACCESS_KEY")
	s3secretKey := os.Getenv("MINIO_SECRET_KEY")

	endpoint :=  s3host + ":" + s3port
	ssl := false

	// content initialization
	minioClient, err := minio.New(endpoint, s3accessKey, s3secretKey, ssl)
	if err != nil {
		Error = 1 //it is not clear at what errors this code works
		ErrorDescription = err.Error()
		return Error , ErrorDescription , s3urllink
	}

	// Creating a new package, if there is then write files to it
	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			Error = 2
			ErrorDescription = err.Error()
			return Error , ErrorDescription , s3urllink

		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// File download ------------------------
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType:contentType})
	if err != nil {
		Error = 3
		ErrorDescription = err.Error()
		return Error , ErrorDescription , s3urllink
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	// Getting a reference to a loaded object ----------
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""  + objectName + "\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Second * 24 * 60 * 60, reqParams)
	if err != nil {
		fmt.Println(err)
		Error = 4
		ErrorDescription = err.Error()
		return Error , ErrorDescription , s3urllink
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
	s3urllink = presignedURL.String()

	return Error , ErrorDescription , s3urllink
}

/* File content .env

	S3_HOST=127.0.0.1
	S3_PORT=9000
	MINIO_ACCESS_KEY=AAAABBBBCCCC
	MINIO_SECRET_KEY=ZZZZXXXXCCCCZZZZXXXXCCCC

*/
