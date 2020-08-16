package oss_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
	"testing"
)

const (
	endpoint         = "http://oss.ceph.cschain.net"
	accessKeyId      = "E6VU593OO63ZIB3D4SCS"
	accessKeySercret = "vWnkPipUWmytfsQD85pwVacwBauDHlP9EZq8qJGJ"
	bucketName       = "cschain-dev"
)

func TestListBucket(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyId, accessKeySercret, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(endpoints.AwsPartitionID),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(false),
	})

	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		HandleError(err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

func TestListObject(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyId, accessKeySercret, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(endpoints.AwsPartitionID),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})

	svc := s3.New(sess)
	params := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	resp, err := svc.ListObjectsV2(params)

	if err != nil {
		HandleError(err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}

func TestCreateBucket(t *testing.T) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKeyId, accessKeySercret, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String(endpoints.AwsPartitionID),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})

	svc := s3.New(sess)

	params := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	_, err = svc.CreateBucket(params)

	if err != nil {
		HandleError(err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucketName)

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		HandleError(err)
	}

	fmt.Printf("Bucket %q successfully created\n", bucketName)
}

func HandleError(err error) {
	fmt.Print("Error:", err)
	os.Exit(-1)
}
