package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment variables")
	}

	fileDir := flag.String("file", "", "path/to/your/file")
	flag.Parse()

	// var key, secret, bucket string
	// if key, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); ok {
	// 	log.Fatal("Failed to load AWS credentials")
	// }

	// if secret, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); ok {
	// 	log.Fatal("Failed to load AWS credentials")
	// }

	// Define the S3 bucket name and the file to upload
	// if bucket, ok := os.LookupEnv("S3_BUCKET"); ok {
	// 	log.Fatal("Failed to load AWS bucket")
	// }
	bucket, ok := os.LookupEnv("S3_BUCKET")
	if !ok {
		log.Fatal("Failed to load AWS bucket")
	}

	region, ok := os.LookupEnv("S3_REGION")
	if !ok {
		log.Fatal("Failed to load AWS bucket")
	}

	key := filepath.Base(*fileDir)

	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an S3 service client
	svc := s3.NewFromConfig(cfg)

	// Open the file for use
	file, err := os.Open(*fileDir)
	if err != nil {
		log.Fatalf("failed to open file %q, %v", *fileDir, err)
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

	fmt.Printf("Successfully uploaded %q to %q\n", *fileDir, bucket)
}
