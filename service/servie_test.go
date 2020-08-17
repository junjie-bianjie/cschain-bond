package service_test

import (
	"cschain-bond/service"
	"testing"
)

func TestUploadByNFT(t *testing.T) {
	service.UploadByNFT()
}

func TestDataCollation(t *testing.T) {
	service.DataCollation()
}

func TestUploadEncryptFile(t *testing.T) {
	service.UploadEncryptFile()
}

func TestReEncryptFile(t *testing.T) {
	service.ReEncryptFile()
}

func TestPreDecryptFile(t *testing.T) {
	service.PreDecryptFile()
}

func TestTemp(t *testing.T) {
	service.UploadEncryptFile()
	service.ReEncryptFile()
	service.PreDecryptFile()
}

func TestTemp2(t *testing.T) {
	service.TempTest()
}
