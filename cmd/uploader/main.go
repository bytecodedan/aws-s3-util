package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/bytecodedan/aws-s3-util/internal/core"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment variables")
	}

	fileDir := flag.String("file", "", "path/to/your/file")
	flag.Parse()

	bucket, ok := os.LookupEnv("S3_BUCKET")
	if !ok {
		log.Fatal("Failed to load AWS bucket")
	}

	region, ok := os.LookupEnv("S3_REGION")
	if !ok {
		log.Fatal("Failed to load AWS bucket")
	}

	key := filepath.Base(*fileDir)

	core.Upload(key, *fileDir, bucket, region)
}
