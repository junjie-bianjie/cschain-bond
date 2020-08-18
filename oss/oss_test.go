package oss_test

import (
	"cschain-bond/oss"
	"fmt"
	"testing"
)

func TestListBucket(t *testing.T) {
	oss.ListBucket()
}

func TestListObject(t *testing.T) {
	oss.ListObjects()
}

func TestUploadFile(t *testing.T) {
	filename := "/Users/bianjie/go/src/cschain-bond/scripts/schema.sql"
	err := oss.UploadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("Successfully uploaded %q\n", filename))
}

func TestDownloadFile(t *testing.T) {
	filename := "/Users/bianjie/go/src/cschain-bond/scripts/schema.sql"
	err := oss.DownloadFile(filename, filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("Successfully download %q\n", filename))
}

func TestDeleteFile(t *testing.T) {
	filename := "0817reEncrypt.text"
	err := oss.DeleteFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fmt.Sprintf("Successfully delete %q\n", filename))
}
