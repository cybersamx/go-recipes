package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	region = "us-east-1"
	bucket = "snowflake-workshop-lab"
	prefix = ""
)

var units = [...]string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB"}

func normalizeSize(size int64) (float64, string) {
	fsize := float64(size)

	step := 0
	for fsize < 0.0 || fsize >= 1000.0 {
		fsize /= 1000.0
		step++
	}

	return fsize, units[step]
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.AnonymousCredentials,
	})
	if err != nil {
		panic(err)
	}

	client := s3.New(sess)

	numPages := 0
	numPrefixes := 0
	numObjects := 0
	size := int64(0)

	// Since we are just cloning `aws s3 ls` function, we just want to get the content for
	// just the level denoted by the prefix. S3 is a flat object store and uses a delimiter
	// and prefixes to emulate folders. So for us to just retrieve the objects for only
	// one level, we set the delimiter and prefix in the inputs.
	inputs := s3.ListObjectsV2Input{
		Bucket:    aws.String(bucket),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String("/"), // If delimiter is unset, we list all objects for the bucket.
	}

	err = client.ListObjectsV2Pages(&inputs, func(page *s3.ListObjectsV2Output, isLastPage bool) bool {
		numPages++
		numObjects += len(page.Contents)

		fmt.Printf("List of s3://%s/%s", bucket, prefix)

		for _, p := range page.CommonPrefixes {
			numPrefixes++
			fmt.Printf("                            PRE %s\n", *p.Prefix)
		}

		for _, obj := range page.Contents {
			size += *obj.Size
			numObjects++

			fsize, funit := normalizeSize(*obj.Size)
			fmt.Printf("%s %9.2f%s %s\n",
				obj.LastModified.Format("2006-01-02 15:04:05"),
				fsize,
				funit,
				*obj.Key,
			)
		}

		return true
	})
	if err != nil {
		panic(err)
	}

	tsize, tunit := normalizeSize(size)
	fmt.Println()
	fmt.Printf("Total size: %.2f%s\n", tsize, tunit)
	fmt.Printf("Total prefixes: %d\n", numPrefixes)
	fmt.Printf("Total objects: %d\n", numObjects)
}
