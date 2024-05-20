S3_UPLOADER=s3-uploader

run_s3_upload: build_s3_uploader
	/bin/$(S3_UPLOADER)

build_s3_uploader:
	go build -o ./bin/$(S3_UPLOADER) ./cmd/uploader/main.go