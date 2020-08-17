package oss

import (
	"cschain-bond/logger"
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
		Region:           aws.String(endpoints.CnNorth1RegionID),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false),
	})

	if err != nil {
		logger.Error("connection s3 failed", logger.String("err", err.Error()))
		panic(err)
	}
}

func ListBucket() {
	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to list buckets, %v", err))
		return
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
		logger.Error(fmt.Sprintf("Unable to list items in bucket %q, %v", bucket, err))
		return
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
}

func UploadFile(originFilePath, desFilename string) error {
	file, err := os.Open(originFilePath)
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to open file %q, %v", originFilePath, err))
		return err
	}
	defer file.Close()

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		// Can also use the `filepath` standard library package to modify the
		// filename as need for an S3 object key. Such as turning absolute path
		// to a relative path.
		Key: aws.String(desFilename),
		// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
		// will be able to optimize memory when uploading large content. io.Reader
		// is supported, but will require buffering of the reader's bytes for
		// each part.
		Body: file,
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to upload %q to %q, %v\n", desFilename, bucket, err))
		return err
	}

	logger.Info(fmt.Sprintf("Successfully uploaded %q to %q\n", desFilename, bucket))
	return nil
}

func DownloadFile(item, localFilePath string) error {
	file, err := os.Create(localFilePath)
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to open file %q, %v", item, err))
		return err
	}
	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to download item %q, %v\n", item, err))
		return err
	}

	logger.Info(fmt.Sprint("Downloaded ", file.Name(), numBytes, "bytes"))
	return nil
}

func DeleteFile(obj string) error {
	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to delete object %q from bucket %q, %v\n", obj, bucket, err))
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(obj),
	})
	if err != nil {
		logger.Error(fmt.Sprintf("Error occurred while waiting for object %q to be deleted, %v\n", obj, err))
		return err
	}

	logger.Info(fmt.Sprintf("Object %q successfully deleted\n", obj))
	return nil
}
