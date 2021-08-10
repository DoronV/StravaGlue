package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestUploadingFiles(t *testing.T) {
	tempFile, err := ioutil.TempFile("tempDir", "tempFile")
	if err != nil {
		t.Errorf("an error occured: %v", err)
	}

	bodyReader, _ := os.Open("tempFile")

	uploadRequest, _ := http.NewRequest("POST", "/combine", bodyReader)
	uploadRequest.Header.Set("Content-Type", "multipart/form-data")
	uploadFiles(uploadRequest)
	os.Remove(tempFile.Name())
	// TODO: create multipart/form to test with
}
