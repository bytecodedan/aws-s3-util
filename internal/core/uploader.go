package core

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func Upload(awsKey, awsSecret, key, fileDir, bucket, region string) {
	// Load the AWS configuration
	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	creds := credentials.NewStaticCredentialsProvider(awsKey, awsSecret, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(region))
	if err != nil {
		log.Fatalf("Unable to load SDK config, %v", err)
	}

	// Create an S3 service client
	svc := s3.NewFromConfig(cfg)

	log.Print(fileDir)
	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		log.Fatalf("Failed to open file %q, %v", fileDir, err)
	}
	defer file.Close()

	// Get the file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Upload the file to S3
	input := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(key),
		Body:          file,
		ContentLength: &size,
		ContentType:   aws.String("application/octet-stream"),
		ACL:           types.ObjectCannedACLPublicRead,
	}

	_, err = svc.PutObject(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to upload file, %v", err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", fileDir, bucket)
}
