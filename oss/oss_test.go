package oss_test

import (
	"github.com/stduolc/aliyun/oss"
	// "log"
	"os"
	"testing"
)

func TestPutObject(t *testing.T) {
	bucketName := "stduolc818"
	objName := "test.mp4"

	uploadFilePath := "d:/temp/2.mp4"

	file, _ := os.Open(uploadFilePath)

	client := oss.NewOSSClient("GhS53sqWHl8riOhH", "dflue7j4z92WBPZQbCfexofFQBj9uP")

	// meta := oss.NewObjectMetadata()
	// info, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// meta.SetContentLength(info.Size())
	// meta.SetContentType("video/mp4")
	// meta.SetContentMd5(file)

	// objectRequest := client.PutObjectRequest(bucketName, objName, file, meta)
	objectRequest := client.PutObjectRequest(bucketName, objName, file)

	// client.PutObjectRequest(bucketName, key, file)

}
