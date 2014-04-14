/*
 */
package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"github.com/stduolc/aliyun/oss/consts"
	// "io/ioutil"
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
		return resp
	} else {
		log.Panic(err)
		return nil
	}
}

// Bucket APIS

func (o *OSSClient) PutBucket(bucketName string) *Bucket {
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
	log.Println("PutBucket : \n", resp)
	return nil
}

func (o *OSSClient) GetBucket(bucketName string) *Bucket {
	meta := o.ObjectMetadata()
	meta.SetHttpHeader(consts.RFC2616_METHOD, "GET")

	meta.SetBucketName(bucketName)
	meta.SetHttpHeader(consts.RFC2616_HOST, meta.GetBucketName()+"."+meta.OSSClient.Auth.Domain)
	meta.SetContentType("")
	t := time.Now()
	meta.SetDate(t.UTC().Format(consts.RFC1123G))

	o.SignMeta(meta)
	resp := o.execute(meta)
	defer resp.Body.Close()
	d := xml.NewDecoder(resp.Body)
	var b Bucket
	d.Decode(&b)
	log.Println("GetBucket : \n", b)
	return &b
}

func (o *OSSClient) DeleteBucket(bucketName string) {
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
	log.Println("DeleteBucket : \n", resp)
	return
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
