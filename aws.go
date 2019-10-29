package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const bucket = "vinhlh.fm"

func getS3(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func getDownloader(sess *session.Session) *s3manager.Downloader {
	return s3manager.NewDownloader(sess)
}

func getSession() *session.Session {
	sess, _ := session.NewSessionWithOptions(session.Options{
		Profile: "s2",
		Config: aws.Config{
			Region: aws.String("ap-southeast-1"),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	return sess
}

// download function (awsFileKey, filePath)
// awsFileKey = advanced-async-js/xxxxxx.webm
// filePath = download/advanced-async-js/lession1.webm
func download(awsFileKey string, filePath string) {
	downloader := getDownloader(getSession())
	f, err := os.Create(filePath)
	if err != nil {
		log("failed to create file", filePath, err)
		return
	}
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(awsFileKey),
	})
	if err != nil {
		log("failed to download file", err)
		return
	}
	fmt.Printf("file downloaded, %d bytes\n", n)
	return
}
