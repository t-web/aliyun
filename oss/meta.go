package oss

import (
	// "crypto/hmac"
	// "crypto/sha1"
	// "encoding/base64"
	"errors"
	"fmt"
	"github.com/stduolc/aliyun/oss/consts"
	// "log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

type ObjectMetadata struct {
	OSSClient *OSSClient
	Req       *http.Request

	kvpool map[string]string

	httpHeader map[string][]string
	ossHeader  map[string][]string

	AES string
}

// ObjectMetadata Factory Function
// You should get a ObjectMetadata before you did any request
func newObjectMetadata(c *OSSClient) (*ObjectMetadata, error) {
	ret := new(ObjectMetadata)
	ret.OSSClient = c
	ret.Req = nil

	ret.kvpool = make(map[string]string)

	ret.httpHeader = make(map[string][]string)
	ret.ossHeader = make(map[string][]string)

	return ret, nil
}

func (meta *ObjectMetadata) CreateRequest() (*http.Request, error) {
	if meta.Req != nil {
		return nil, errors.New("meta had been used!")
	}
	// meta.signr()
	meta.Req = new(http.Request)

	uri := meta.GetURI()
	meta.Req.URL, _ = url.Parse(uri)

	meta.Req.Header = make(map[string][]string)
	meta.Req.Method = meta.GetMethod()

	//	meta.writeHttpHeaders()

	//	meta.writeOSSHeaders()

	// MARK
	return meta.Req, nil
}

func (obj *ObjectMetadata) SetHttpHeader(key, value string) {
	if v, ok := obj.httpHeader[key]; ok {
		obj.httpHeader[key] = append(v, value)
		sort.Strings(v)
		return
	} else {
		obj.httpHeader[key] = make([]string, 1)
		obj.httpHeader[key][0] = value
		return
	}
}

func (obj *ObjectMetadata) GetHttpHeader(key string) []string {
	if v, ok := obj.httpHeader[key]; ok {
		return v
	} else {
		return []string{""}
	}
}

func (obj *ObjectMetadata) SetOSSHeader(key, value string) {
	if _, ok := obj.ossHeader[key]; ok {
		obj.ossHeader[key] = append(obj.ossHeader[key], value)
		sort.Strings(obj.ossHeader[key])
		return
	} else {
		obj.ossHeader[key] = make([]string, 1)
		obj.ossHeader[key][0] = value
		return
	}
}

func (obj *ObjectMetadata) GetOSSHeader(key string) []string {
	if val, ok := obj.httpHeader[key]; ok {
		return val
	} else {
		return []string{""}
	}
}

// GETTER/SETTER FOR HTTP HEADER

func (obj *ObjectMetadata) GetContentLength() int64 {
	if ret, ok := obj.httpHeader[consts.RFC2616_CONTENT_LENGTH]; ok {
		i, _ := strconv.ParseInt(ret[0], 10, 64)
		return i
	} else {
		return int64(-1)
	}
}

func (obj *ObjectMetadata) SetContentLength(i int64) {
	if v, ok := obj.httpHeader[consts.CONTENT_LENGTH]; ok {
		v[0] = strconv.FormatInt(i, 10)
		return
	} else {
		obj.httpHeader[consts.RFC2616_CONTENT_LENGTH] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_CONTENT_LENGTH][0] = strconv.FormatInt(i, 10)
		return
	}
}

func (obj *ObjectMetadata) GetContentType() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_CONTENT_TYPE]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetContentType(t string) {
	if _, ok := obj.httpHeader[consts.RFC2616_CONTENT_TYPE]; ok {
		obj.httpHeader[consts.RFC2616_CONTENT_TYPE][0] = t
		return
	} else {
		obj.httpHeader[consts.RFC2616_CONTENT_TYPE] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_CONTENT_TYPE][0] = t
		return
	}
}

func (obj *ObjectMetadata) GetContentEncoding() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_CONTENT_ENCODING]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetContentEncoding(t string) {
	if _, ok := obj.httpHeader[consts.RFC2616_CONTENT_ENCODING]; ok {
		obj.httpHeader[consts.RFC2616_CONTENT_ENCODING][0] = t
		return
	} else {
		obj.httpHeader[consts.RFC2616_CONTENT_ENCODING] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_CONTENT_ENCODING][0] = t
		return
	}
}

func (obj *ObjectMetadata) GetCacheControl() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_CACHE_CONTROL]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetCacheControl(c string) {
	if _, ok := obj.httpHeader[consts.RFC2616_CACHE_CONTROL]; ok {
		obj.httpHeader[consts.RFC2616_CACHE_CONTROL][0] = c
		return
	} else {
		obj.httpHeader[consts.RFC2616_CACHE_CONTROL] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_CACHE_CONTROL][0] = c
		return
	}
}

func (obj *ObjectMetadata) GetContentMd5() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_CONTENT_MD5]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetContentMd5(c string) {
	if _, ok := obj.httpHeader[consts.RFC2616_CONTENT_MD5]; ok {
		obj.httpHeader[consts.RFC2616_CONTENT_MD5][0] = c
		return
	} else {
		obj.httpHeader[consts.RFC2616_CONTENT_MD5] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_CONTENT_MD5][0] = c
		return
	}
}

func (obj *ObjectMetadata) GetContentDisposition() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_CONTENT_MD5]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetContentDisposition(c string) {
	if _, ok := obj.httpHeader[consts.RFC2616_CONTENT_DISPOSITION]; ok {
		obj.httpHeader[consts.RFC2616_CONTENT_DISPOSITION][0] = c
		return
	} else {
		obj.httpHeader[consts.RFC2616_CONTENT_DISPOSITION] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_CONTENT_DISPOSITION][0] = c
		return
	}
}

func (obj *ObjectMetadata) GetDate() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_DATE]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetDate(date string) {
	if _, ok := obj.httpHeader[consts.RFC2616_DATE]; ok {
		obj.httpHeader[consts.RFC2616_DATE][0] = date
		return
	} else {
		obj.httpHeader[consts.RFC2616_DATE] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_DATE][0] = date
		return
	}
}

func (obj *ObjectMetadata) GetLastModified() string {
	if ret, ok := obj.httpHeader[consts.RFC2616_LAST_MODIFIED]; ok {
		return ret[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetLastModified(s string) {
	if _, ok := obj.httpHeader[consts.RFC2616_LAST_MODIFIED]; ok {
		obj.httpHeader[consts.RFC2616_LAST_MODIFIED][0] = s
		return
	} else {
		obj.httpHeader[consts.RFC2616_LAST_MODIFIED] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_LAST_MODIFIED][0] = s
		return
	}
}

// GETTER/SETTER FOR URL

func (obj *ObjectMetadata) GetURI() string {
	uri := obj.GetScheme() + "://"
	uri = uri + obj.GetHost()
	return uri
}

func (obj *ObjectMetadata) GetScheme() string {
	return obj.OSSClient.Auth.Scheme
}

func (m *ObjectMetadata) GetHost() string {
	if _, ok := m.httpHeader[consts.HOST]; ok {
		return m.httpHeader[consts.RFC2616_HOST][0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) GetBucketName() string {
	v, ok := obj.kvpool["BucketName"]
	if ok {
		return v
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetBucketName(s string) {
	if s == "" {
		delete(obj.kvpool, "BucketName")
	} else {
		obj.kvpool["BucketName"] = s
	}
}

func (obj *ObjectMetadata) GetObjectName() string {
	v, ok := obj.kvpool["ObjectName"]
	if ok {
		return v
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetObjectName(s string) {
	if s == "" {
		delete(obj.kvpool, "ObjectName")
	} else {
		obj.kvpool["ObjectName"] = s
	}
}

// GETTER/SETTER FOR Canonicalize Resources
func (obj *ObjectMetadata) GetCanonicalizedResources() string {
	ret := ""
	ret = fmt.Sprint(ret, "/", obj.GetBucketName(), "/", obj.GetObjectName())
	return ret
}

// 缺少子资源
func (obj *ObjectMetadata) GetResources() string {
	ret := ""
	objName := obj.GetObjectName()
	ret = ret + "/" + objName
	return ret
}

func (obj *ObjectMetadata) GetServerSideEncryption() []string {
	ret, ok := obj.httpHeader[consts.SERVER_SIDE_ENCRYPTION]
	if ok {
		return ret
	} else {
		return []string{""}
	}
}

func (obj *ObjectMetadata) SetServerSideEncryption(s string) {
	if _, ok := obj.httpHeader[consts.SERVER_SIDE_ENCRYPTION]; ok {
		obj.httpHeader[consts.SERVER_SIDE_ENCRYPTION][0] = s
		return
	} else {
		obj.httpHeader[consts.SERVER_SIDE_ENCRYPTION] = make([]string, 1)
		obj.httpHeader[consts.SERVER_SIDE_ENCRYPTION][0] = s
		return
	}
}

func (m *ObjectMetadata) GetMethod() string {
	if v, ok := m.httpHeader[consts.RFC2616_METHOD]; ok {
		return v[0]
	} else {
		return ""
	}
}

func (obj *ObjectMetadata) SetMethod(method string) {
	if _, ok := obj.httpHeader[consts.RFC2616_METHOD]; ok {
		obj.httpHeader[consts.RFC2616_METHOD][0] = method
		return
	} else {
		obj.httpHeader[consts.RFC2616_METHOD] = make([]string, 1)
		obj.httpHeader[consts.RFC2616_METHOD][0] = method
		return
	}
}

func (m *ObjectMetadata) SetHost(host string) {
	allHost := ""
	allHost = m.GetBucketName() + "." + host
	m.httpHeader[consts.RFC2616_HOST][0] = allHost
}

func (m *ObjectMetadata) WriteHttpHeaders() {
	for k, v := range m.httpHeader {
		m.Req.Header.Set(k, v[0])
	}
}
