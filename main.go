package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	var (
		output string
	)
	flag.StringVar(&output, "output", "/dev/stdout", "Place to send the output")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-output=FILEPATH] <bucket> <key>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	bucket, key := flag.Arg(0), flag.Arg(1)

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("s4cat: unable to load config: %v", err)
	}
	cfg.Region = "eu-west-1"
	svc := s3.NewFromConfig(cfg)

	resp, err := svc.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Fatalf("GetObject Failed: %#+v", err)
	}
	fd, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	n, err := io.Copy(fd, resp.Body)
	if err != nil {
		log.Fatal("Copy failed after", n, "bytes:", err)
	}
}
