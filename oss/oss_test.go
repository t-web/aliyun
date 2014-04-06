package oss_test

import (
	"github.com/stduolc/aliyun/oss"
	"github.com/stduolc/aliyun/oss/consts"
	// "log"
	// "net/http"
	// "net/url"
	// "time"

	// "os"
	"testing"
)

func testPutObject(t *testing.T) {
	// bucketName := "stduolc818"
	// objName := "test.mp4"

	// uploadFilePath := "d:/temp/2.mp4"

	// file, _ := os.Open(uploadFilePath)

	// client := oss.NewOSSClient("GhS53sqWHl8riOhH", "dflue7j4z92WBPZQbCfexofFQBj9uP")

	// meta := oss.NewObjectMetadata()
	// info, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// meta.SetContentLength(info.Size())
	// meta.SetContentType("video/mp4")
	// meta.SetContentMd5(file)

	// objectRequest := client.PutObjectRequest(bucketName, objName, file, meta)
	// objectRequest := client.PutObjectRequest(bucketName, objName, file)

	// client.PutObjectRequest(bucketName, key, file)

}

// Test for SignMeta
func Test001(t *testing.T) {
	o := oss.NewOSSClient("GhS53sqWHl8riOhH", "dflue7j4z92WBPZQbCfexofFQBj9uP", consts.DOMAIN_DEFAULT)

	meta := o.ObjectMetadata()

	meta.SetHttpHeader(consts.RFC2616_METHOD, "PUT")

	meta.SetBucketName("stduolc919")
	meta.SetHttpHeader(consts.RFC2616_HOST, meta.GetBucketName()+"."+consts.DOMAIN_DEFAULT)
	meta.SetContentType("")
	meta.SetDate("Thu, 03 Apr 2014 07:06:42 GMT")

	// right Authorization is : "OSS GhS53sqWHl8riOhH:zFBtZV19pyBqTt5yuBEdQOI0kSg="
	o.SignMeta(meta)
}

// Test for PutBucket
func Test002(t *testing.T) {
	bucketName := "stduolc888"
	host := consts.DOMAIN_DEFAULT
	client := oss.NewOSSClient("GhS53sqWHl8riOhH", "dflue7j4z92WBPZQbCfexofFQBj9uP", consts.DOMAIN_DEFAULT)
	client.PutBucket(bucketName, host)
}
