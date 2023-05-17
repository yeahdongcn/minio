package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	var sess *session.Session
	var err error

	awsConfig := aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String("TODO"),
		Endpoint:         aws.String("TODO"),
		Credentials:      credentials.NewStaticCredentials("TODO", "TODO", ""),
	}

	sess, err = session.NewSession(&awsConfig)
	if err != nil {
		panic(err)
	}

	sessionClient := s3.New(sess)
	response, err := sessionClient.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String("TODO"),
		Prefix: aws.String("TODO"),
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(response)
}
