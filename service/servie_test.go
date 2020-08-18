package service_test

import (
	"cschain-bond/service"
	"testing"
)

func TestUploadByNFT(t *testing.T) {
	service.UploadByNFT()
}

func TestDataCollation(t *testing.T) {
	service.DataCollation("yoeu")
}

func TestUploadEncryptFile(t *testing.T) {
	err := service.UploadEncryptFile("../scripts/origin.text")
	if err != nil {
		t.Fatal(err)
	}
}

func TestReEncryptFile(t *testing.T) {
	err := service.ReEncryptFile("origin.text", "origin_encrypt.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPreDecryptFile(t *testing.T) {
	_, err := service.PreDecryptFile("0817reEncrypt.text")
	if err != nil {
		t.Fatal(err)
	}
}

func TestOnce(t *testing.T) {
	err := service.UploadEncryptFile("../scripts/origin.text")
	if err != nil {
		t.Fatal(err)
	}

	err = service.ReEncryptFile("origin.text", "origin_encrypt.txt")
	if err != nil {
		t.Fatal(err)
	}

	res, err := service.PreDecryptFile("origin_encrypt.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestTwo(t *testing.T) {
	service.GrpClient()
}
