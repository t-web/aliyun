/*

*/
package oss

import (
	"encoding/base64"
	"fmt"
	"github.com/stduolc/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	//	"strconv"
)

const (
	debug       = false
	DefaultHost = "http://oss.aliyuncs.com"

	HangZhou         = "oss-cn-hangzhou.aliyuncs.com"
	QingDao          = "oss-cn-qingdao.aliyuncs.com"
	HangZhouInternal = "oss-cn-hangzhou-internal.aliyuncs.com"
	QingdaoInternal  = "oss-cn-qingdao-internal.aliyuncs.com"
	DefaultRegion    = "oss.aliyuncs.com"
)

type OSSClient struct {
	Auth   *Auth
	Client *http.Client
}

func NewOSSClient(accessId, accessKey string) *OSSClient {
	ret := new(OSSClient)
	ret.Auth = &Auth{AccessKey: accessId, SecretKey: accessKey}
	ret.Client = new(http.Client)
	return ret
}

//MD5计算上存在问题，需要把整个文件读入内存。
func (c *OSSClient) PutObjectRequest(bucketName string, objName string, reader io.Reader, meta *ObjectMetadata) {
	req := http.Request{}
	req.Header = make(map[string][]string)

	// 将meta的基本信息同步到Req
	meta.PushRequest(&req)

	// 根据不同方法改写定制Req
	req.Method = "PUT"
	req.URL = &url.URL{}
	req.URL.Scheme = "http"
	req.URL.Host = fmt.Sprint(bucketName, ".", HangZhou)
	req.Header.Set("Host", req.URL.Host)
	now := time.Now()
	req.Header.Set("Date", now.Format("Mon, 2 Jan 2006 15:04:05 GMT"))
	req.URL.Path = fmt.Sprint("/", objName)

	// 将Req中的特殊信息改写meta
	meta.MergeRequest(&req)

	_, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	//req.Write(bytes)
	resp, err := c.Client.Do(&req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("\n:REQUEST:\n", req.Header)
	log.Println("\n:RESPONSE:\n", resp.Header, "\n", string(body))
}

func (c *OSSClient) PutObjectRequest(bucketName string, key string, file os.File) (ret *PutObjectRequest, err error) {
	req := http.Request{}
	req.Header = make(map[string][]string)

	// 将meta的基本信息同步到Req
	meta.PushRequest(&req)

	// 根据不同方法改写定制Req
	req.Method = "PUT"
	req.URL = &url.URL{}
	req.URL.Scheme = "http"
	req.URL.Host = fmt.Sprint(bucketName, ".", HangZhou)
	req.Header.Set("Host", req.URL.Host)
	now := time.Now()
	req.Header.Set("Date", now.Format("Mon, 2 Jan 2006 15:04:05 GMT"))
	req.URL.Path = fmt.Sprint("/", objName)

	// 将Req中的特殊信息改写meta
	meta.MergeRequest(&req)

	/*
		, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal(err)
		}
	*/
	//req.Write(bytes)
	resp, err := c.Client.Do(&req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("\n:REQUEST:\n", req.Header)
	log.Println("\n:RESPONSE:\n", resp.Header, "\n", string(body))
}

func (c *OSSClient) PutObject(putReq *PutObjectRequest) (ret *PutObjectResult, err error) {
	if nil == putReq {
		return nil, Error("Nil param putObjectRequest")
	}
	bucketName := putReq.BucketName
	key := putReq.Key
	metadata := putReq.Metadata
	inputStream := putReq.InputStream

	if metadata == nil {
		metadata = NewObjectMetadata()
	}

	if bucketName == nil || bucketName == "" {
		return nil, Error("The bucket name parameter must be specified when uploading an object")
	}

	if key == nil || key == "" {
		return nil, Error("The key parameter must be specified when uploading an object")
	}

	if putReq.File != nil {
		file := putReq.File
		fileInfo := file.Stat()
		metadata.SetContentLength(fileInfo.Size())

		if ctype, ok := metadata.GetContentType(); ok != nil {
			filetype := TypeByFileName(fileInfo.Name())
			metadata.SetContenType(filetype)
		}

		// TODO here is something strange. If the api realy like the document said, maybe we should check the inputstream is not nil.
		// and then read length bytes of data from the inputstream to buff, and calculate the sum5md.
		md5 := util.MD5File(file)
		contentMD5 := base64.StdEncoding.EncodeToString(md5)
		metadata.setContentMd5(contentMD5)

	}

}
