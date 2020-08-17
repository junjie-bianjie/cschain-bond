package oss

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

const (
	endpoint         = "http://oss.ceph.cschain.net"
	accessKeyId      = "E6VU593OO63ZIB3D4SCS"
	accessKeySercret = "vWnkPipUWmytfsQD85pwVacwBauDHlP9EZq8qJGJ"
	bucket           = "cschain-dev/"
)

var sess *session.Session

// Encapsulation oss sdk ,you can refer to https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html
func init() {
	var err error
	sess, err = session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyId, accessKeySercret, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(endpoints.AwsPartitionID),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})

	if err != nil {
		// TODO handle the error
		panic(err)
	}
}

func ListBucket() {
	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		// TODO handle the error
		panic(err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func ListObjects() {
	svc := s3.New(sess)

	params := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	resp, err := svc.ListObjects(params)

	if err != nil {
		// TODO handle the error
		panic(err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}

func UploadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		// TODO handle the error
		panic(err)
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(filename),
		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
	})
	if err != nil {
		// TODO handle the error
		fmt.Printf("Unable to upload %q to %q, %v\n", filename, bucket, err)
		panic(err)
	}

	// TODO use log to print
	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
	return nil
}

func DownloadFile(item string) error {
	file, err := os.Create(item)
	if err != nil {
		// TODO handle the error
		panic(err)
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		// TODO handle the error
		fmt.Printf("Unable to download item %q, %v\n", item, err)
		panic(err)
	}

	// TODO use log to print
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return nil
}

func DeleteFile(obj string) error {
	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		// TODO handle the error
		fmt.Printf("Unable to delete object %q from bucket %q, %v\n", obj, bucket, err)
		panic(err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		// TODO handle the error
		fmt.Printf("Error occurred while waiting for object %q to be deleted, %v\n", obj, err)
		panic(err)
	}

	// TODO use log to print
	fmt.Printf("Object %q successfully deleted\n", obj)
	return nil
}
