package main

import (
	"io"
	"log"
	"os"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/s3"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: s4cat <bucket> <key>")
	}
	bucket, key := os.Args[1], os.Args[2]

	creds := aws.IAMCreds()

	s := s3.New(creds, "eu-west-1", nil)
	resp, err := s.GetObject(&s3.GetObjectRequest{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatalf("GetObject Failed: %#+v", err)
	}
	n, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal("Copy failed after", n, "bytes:", err)
	}
}
