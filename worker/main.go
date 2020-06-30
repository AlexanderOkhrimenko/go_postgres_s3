package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go"
	"go_postgres_s3/worker/modules"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"
)

type JobRow struct {
	id               uint64
	error            sql.NullInt64
	errordescription sql.NullString
	command          sql.NullString
	status           sql.NullString
	complete         sql.NullInt64
	task             sql.NullString
	priority         sql.NullString
	resulturl        sql.NullString
	resultsurl       sql.NullString
	duration         sql.NullFloat64
	outobjects       sql.NullString
}

func main() {

	for {

		fmt.Println("Job search.....")

		id, error, command, status, complete, task, priority, duration, outobjects := SelectMinWaitTask()

		// Show what id took in the work
		fmt.Printf("task %s in work id = %d \n", task, id)

		fmt.Println(" - - - - - - ")
		fmt.Println("error - ", error)
		fmt.Println("command - ", command)
		fmt.Println("status - ", status)
		fmt.Println("complete - ", complete)
		fmt.Println("task - ", task)
		fmt.Println("priority - ", priority)
		fmt.Println("duration - ", duration)
		fmt.Println("outobjects - ", outobjects)
		fmt.Println(" - - - - - - ")

		// Random start time of search in the database, improves the work of several workers
		rand.Seed(time.Now().UTC().UnixNano())
		t := rand.Intn(6)
		fmt.Println("Start worker time = ", t)
		time.Sleep(time.Duration(t) * time.Second)

	}

}

func SelectMinWaitTask() (id uint64, error int, command string, status string, complete int, task string, priority string, duration float64, outobjects string) {

	//var status = "wait" //

	db := modules.ConnectPG()
	defer db.Close(context.Background())

	row := db.QueryRow(context.Background(), "select * from jobs where id = (SELECT MIN(id) FROM jobs WHERE status = 'wait' )")

	var q JobRow

	err := row.Scan(&q.id, &q.error, &q.errordescription, &q.command, &q.status, &q.complete, &q.task, &q.priority, &q.resulturl, &q.resultsurl, &q.duration, &q.outobjects)

	if err != nil {
		//panic(err)
	}

	// Checking the read parameters for validation and forming the final values.
	// Id
	id = q.id

	//  error
	if q.error.Valid {
		error = int(q.error.Int64)
	} else {
		error = 0
	}

	//  command
	if q.command.Valid {
		command = q.command.String
	} else {
		command = ""
	}

	//  status
	if q.status.Valid {
		status = q.status.String
	} else {
		status = ""
	}

	//  complete
	if q.complete.Valid {
		complete = int(q.complete.Int64)
	} else {
		complete = 0
	}

	//  Проверка task
	if q.task.Valid {
		task = q.task.String
	} else {
		task = ""
	}

	//  Проверка priority
	if q.priority.Valid {
		priority = q.priority.String
	} else {
		priority = ""
	}

	// duration
	if q.duration.Valid {
		duration = float64(q.duration.Float64)
	} else {
		duration = 0
	}

	// outobjects
	if q.outobjects.Valid {
		outobjects = q.outobjects.String
	} else {
		outobjects = ""
	}

	return
}

func saveToS3(bucketName string, location string, objectName string, filePath string, contentType string) (Error int, ErrorDescription string, s3urllink string) {

	// If the IP and PORT data are incorrect, an error is generated only after the timeout expires ~ 90 сек

	s3host := os.Getenv("S3_HOST")
	s3port := os.Getenv("S3_PORT")
	s3accessKey := os.Getenv("MINIO_ACCESS_KEY")
	s3secretKey := os.Getenv("MINIO_SECRET_KEY")

	endpoint := s3host + ":" + s3port
	ssl := false

	// content initialization
	minioClient, err := minio.New(endpoint, s3accessKey, s3secretKey, ssl)
	if err != nil {
		Error = 1 //it is not clear at what errors this code works
		ErrorDescription = err.Error()
		return Error, ErrorDescription, s3urllink
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
			return Error, ErrorDescription, s3urllink

		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// File download ------------------------
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		Error = 3
		ErrorDescription = err.Error()
		return Error, ErrorDescription, s3urllink
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, n)

	// Getting a reference to a loaded object ----------
	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	// Generates a presigned url which expires in a day.
	presignedURL, err := minioClient.PresignedGetObject(bucketName, objectName, time.Second*24*60*60, reqParams)
	if err != nil {
		fmt.Println(err)
		Error = 4
		ErrorDescription = err.Error()
		return Error, ErrorDescription, s3urllink
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
	s3urllink = presignedURL.String()

	return Error, ErrorDescription, s3urllink
}
