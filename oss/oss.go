/*
 */
package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/stduolc/aliyun/oss/consts"
	"io/ioutil"
	"log"
	"net/http"
	// "net/url"
	"strings"
	"time"
)

const (
	debug = false
)

type Auth struct {
	AccessKey string
	SecretKey string
	Scheme    string
	Domain    string
}

type OSSClient struct {
	Auth   *Auth
	Client *http.Client
}

func NewOSSClient(accessId, accessKey, domain string) *OSSClient {
	ret := new(OSSClient)
	ret.Auth = &Auth{AccessKey: accessId, SecretKey: accessKey, Scheme: "http", Domain: domain}
	tr := &http.Transport{
		DisableCompression: true,
		DisableKeepAlives:  false,
	}
	ret.Client = &http.Client{Transport: tr}
	return ret
}

func (o *OSSClient) execute(meta *ObjectMetadata) *http.Response {
	o.SignMeta(meta)
	req, err := meta.CreateRequest()

	meta.WriteHttpHeaders()

	if err != nil {
		log.Panic(err)
		return nil
	}

	resp, err := o.Client.Do(req)

	// DEVELOP ACHOR
	if err == nil {
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(bytes))
		return resp
	} else {
		log.Panic(err)
		return nil
	}
}

// Bucket APIS

func (o *OSSClient) PutBucket(bucketName, host string) *Bucket {
	bucketName = strings.Trim(bucketName, " ")

	meta := o.ObjectMetadata()
	meta.SetHttpHeader(consts.RFC2616_METHOD, "PUT")

	meta.SetBucketName(bucketName)
	meta.SetHttpHeader(consts.RFC2616_HOST, meta.GetBucketName()+"."+meta.OSSClient.Auth.Domain)
	meta.SetContentType("")
	t := time.Now()
	meta.SetDate(t.UTC().Format(consts.RFC1123G))

	o.SignMeta(meta)

	resp := o.execute(meta)
	log.Println("dd ", resp)
	return nil
}

func (o *OSSClient) DeleteBucket(bucketName, host string) {
	bucketName = strings.Trim(bucketName, " ")

	meta := o.ObjectMetadata()
	meta.SetHttpHeader(consts.RFC2616_METHOD, "DELETE")

	meta.SetBucketName(bucketName)
	meta.SetHttpHeader(consts.RFC2616_HOST, meta.GetBucketName()+"."+meta.OSSClient.Auth.Domain)
	meta.SetContentType("")
	t := time.Now()
	meta.SetDate(t.UTC().Format(consts.RFC1123G))

	o.SignMeta(meta)

	resp := o.execute(meta)
	log.Println("dd ", resp)
	return
}

func (o *OSSClient) GetBucket(bucketname, host string) {

}

//

func (c *OSSClient) ObjectMetadata() *ObjectMetadata {
	ret, err := newObjectMetadata(c)
	if err != nil {
		log.Println(err)
		return nil
	}
	ret.SetHttpHeader(consts.RFC2616_USER_ANGENT, "aliyun-sdk-go/0.1")
	ret.SetHttpHeader(consts.RFC2616_CONNECTION, "Keep-Alive")
	return ret
}

func (o *OSSClient) SignMeta(meta *ObjectMetadata) {
	method := meta.GetHttpHeader(consts.RFC2616_METHOD)[0]
	contentMd5 := meta.GetHttpHeader(consts.RFC2616_CONTENT_MD5)[0]
	contentType := meta.GetHttpHeader(consts.RFC2616_CONTENT_TYPE)[0]
	date := meta.GetHttpHeader(consts.RFC2616_DATE)[0]
	str := fmt.Sprint(method, "\n", contentMd5, "\n", contentType, "\n", date, "\n")

	// add CanonicalizedResources
	res := meta.GetCanonicalizedResources()
	str = fmt.Sprint(str, res)

	mac := hmac.New(sha1.New, []byte(o.Auth.SecretKey))
	mac.Write([]byte(str))
	dst := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	signature := dst

	authorization := fmt.Sprint("OSS ", o.Auth.AccessKey, ":", signature)
	meta.SetHttpHeader(consts.RFC2616_AUTHORIZATION, authorization)
}

//MD5计算上存在问题，需要把整个文件读入内存。
/*
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
*/

/*
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

//		, err := ioutil.ReadAll(reader)
//		if err != nil {
//		 	log.Fatal(err)
//		 }

	//req.Write(bytes)
	resp, err := c.Client.Do(&req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("\n:REQUEST:\n", req.Header)
	log.Println("\n:RESPONSE:\n", resp.Header, "\n", string(body))
}
*/

/*
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
		metadata.setcontentmd5(contentmd5)

		req := createrequest(bucketname, key, putreq, method.put)

	}
}
*/
