package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Auth struct {
	AccessKey string
	SecretKey string
}

type ObjectMetadata struct {
	auth      *Auth
	ossClient *OSSClient
	req       *http.Request
	method    string

	host          string
	contentLength int64
	contentMD5    string
	contentType   string
	date          string
	resource      string

	signature     string
	authorization string

	// the date when the object is no longer cacheable
	httpExpiresDate time.Date

	// The time this object will expire and be completely removed from S3, or null if this object will never expire.
	// This and the expiration time rule aren't stored in the metadata map because the header contains both the time and the rule.
	expirationTime time.Date

	userMetadata map[string]string
	metadata     map[string]string
	AES          string
}

// ObjectMetadata Factory Function
// You should get a ObjectMetadata before you did any request
func NewObjectMetadata() *ObjectMetadata {
	ret := new(ObjectMetadata)
	initMetadata(ret)
	return ret
}

func (obj *ObjectMetadata) GetUserMetadata() map[string]string {
	return obj.userMetadata
}

// Sets the custom user-metadata for the associated object.
func (obj *ObjectMetadata) SetUserMetadata(usermetadata map[string]string) {
	obj.userMetadata = usermetadata
}

func (obj *ObjectMetadata) SetHeader(key, value string) {
	obj.metadata[key] = value
}

func (obj *ObjectMetadata) AddUserMetadata(key, value string) {
	obj.userMetadata[key] = value
}

func (obj *ObjectMetadata) GetLastModified() (time.Date, error) {
	date, ok := obj.metadata[Header.LAST_MODIFIED]
	if ok == nil {
		ret := time.Parse(TIME_LAYOUT, date)
		return ret, nil
	} else {
		return nil, Error("no last_modified header")
	}
}

func (obj *ObjectMetadata) SetLastModified(date time.Date) {
	obj.metadata[Header.LAST_MODIFIED] = string(date)
}

func (obj *ObjectMetadata) GetContentLength() (int64, error) {
	ret, ok := obj.metadata[Header.CONTENT_LENGTH]
	if ok == nil {
		return strconv.ParseInt(ret, 10, 64), nil
	} else {
		return nil, Error("no Content-length header")
	}
}

func (obj *ObjectMetadata) SetContentLength(i int64) {
	obj.metadata[Header.CONTENT_LENGTH] = strconv.FormatInt(i, 10)
}

func (obj *ObjectMetadata) GetContentType() (string, error) {
	ret, ok := obj.metadata[Header.CONTENT_TYPE]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no Content-Type header")
	}
}

func (obj *ObjectMetadata) SetContenType(t string) {
	obj.metadata[Header.CONTENT_TYPE] = t
}

func (obj *ObjectMetadata) GetContentEncoding() (string, error) {
	ret, ok := obj.metadata[Header.CONTENT_ENCODING]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no Content-Encoding header")
	}
}

func (obj *ObjectMetadata) SetContentEncoding(e string) {
	obj.metadata[Header.CONTENT_ENCODING] = e
}

func (obj *ObjectMetadata) GetCacheControl() (string, error) {
	ret, ok := obj.metadata[Header.CACHE_CONTROL]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no " + Header.CACHE_CONTROL + " header")
	}
}

func (obj *ObjectMetadata) SetCacheControl(c string) {
	obj.metadata[Header.CACHE_CONTROL] = c
}

func (obj *ObjectMetadata) GetContentMd5() (string, error) {
	ret, ok := obj.metadata[Header.CONTENT_MD5]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no " + Header.CONTENT_MD5 + " header")
	}
}

func (obj *ObjectMetadata) SetContentMd5(md5Base64 string) {
	if md5Base64 == nil {
		delete(obj.metadata, Header.CONTENT_MD5)
	} else {
		obj.metadata[Header.CONTENT_MD5] = md5Base64
	}
}

func (obj *ObjectMetadata) GetContentDisposition() (string, error) {
	ret, ok := obj.metadata[Header.CONTENT_DISPOSITION]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no " + Header.CONTENT_DISPOSITION + " header")
	}
}

func (obj *ObjectMetadata) SetContentDisposition(d string) {
	obj.metadata[Header.CONTENT_DISPOSITION] = d
}

func (obj *ObjectMetadata) GetETage() (string, error) {
	ret, ok := obj.metadata[Header.ETAG]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no " + Header.ETAG + " header")
	}
}

func (obj *ObjectMetadata) GetServerSideEncryption() (string, error) {
	ret, ok := obj.metadata[Header.SERVER_SIDE_ENCRYPTION]
	if ok == nil {
		return ret, nil
	} else {
		return nil, Error("no " + Header.SERVER_SIDE_ENCRYPTION + " header")
	}
}

func (obj *ObjectMetadata) SetServerSideEncryption(s string) {
	obj.metadata[Header.SERVER_SIDE_ENCRYPTION] = s
}

func (obj *ObjectMetadata) GetExpirationTime() time.Date {
	return obj.expirationTime
}

func (obj *ObjectMetadata) SetExpirationTime(t time.Date) {
	obj.expirationTime = t
}

func (obj *ObjectMetadata) GetHttpExpiresDate() time.Date {
	return obj.httpExpiresDate
}

func (obj *ObjectMetadata) SetHttpExpiresDate(t time.Date) {
	obj.httpExpiresDate = t
}

//
func (m *ObjectMetadata) setRequest(req *http.Request) {
	m.setContentType(req)
	m.setDate(req)
	m.setOSSHeaders(req)
	m.setResource(req)

	m.signr()
	m.req = req
}

func (m *ObjectMetadata) PushRequest(req *http.Request) {
	req.Header.Set("Content-Length", strconv.FormatInt(m.GetContentLength(), 10))
	req.Header.Set("Content-Type", m.GetContentType())
	req.Header.Set("Date", m.date)
	req.Header.Set("Authorization", m.authorization)
	m.setRequest(req)
}

func (m *ObjectMetadata) MergeRequest(req *http.Request) {
	m.method = req.Method
	m.host = req.URL.Host
}

func (m *ObjectMetadata) GetMethod() string {
	return m.method
}

// Learned from S3 sdk
func (m *ObjectMetadata) setOSSHeaders(req *http.Request) {
	metas := make(UserMetadata)
	for k, v := range req.Header {
		key := strings.ToLower(k)
		tmp := ""
		leng := len(v)
		for i := 0; i < leng; i++ {
			if i == 0 {
				tmp = v[i]
				continue
			}
			tmp = tmp + ","
			tmp = tmp + v[i]
		}
		val := strings.ToLower(tmp)
		if strings.HasPrefix(key, "x-oss-") {
			//metas = append(metas, {key, val})
			metas[key] = val
		}
	}
	sort.Sort(metas)
	m.userMetadata = metas
}

func (m *ObjectMetadata) setResource(req *http.Request) {
	tmp := req.Header.Get("Host")
	sp := strings.Split(tmp, ".")
	bucket := sp[0]
	name := req.URL.Path
	name = fmt.Sprint("/", bucket, name)
	m.resource = name
}

func (m *ObjectMetadata) signr() {
	str := fmt.Sprint(m.method, "\n", m.contentMD5, "\n", m.contentType, "\n", m.date, "\n")
	for k, v := range m.userMetadata {
		str = fmt.Sprint(str, k, ":", v, "\n")
	}
	str = fmt.Sprint(str, m.resource)
	fmt.Println(str)
	mac := hmac.New(sha1.New, []byte(m.auth.SecretKey))
	mac.Write([]byte(str))
	dst := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	m.signature = dst

	m.authorization = fmt.Sprint("OSS ", m.auth.AccessKey, ":", m.signature)
	fmt.Println(dst)
}

func initMetadata(m *ObjectMetadata) {
	m.auth = nil
	m.ossClient = nil
	m.host = ""
}

// UserMetadata Method
func keyset(m map[string]string) []string {
	ret := make([]string, len(m), len(m))
	i := 0
	for k, _ := range m {
		ret[i] = k
		i++
	}
	return ret
}

func (h UserMetadata) Itok(i int) string {
	kset := keyset(h)
	return kset[i]
}

func (h UserMetadata) Len() int {
	return len(h)
}

func (h UserMetadata) Swap(i, j int) {
	ik := h.Itok(i)
	jk := h.Itok(j)
	h[ik], h[jk] = h[jk], h[ik]
}

func (h UserMetadata) Less(i, j int) bool {
	ret := true

	ikey := h.Itok(i)
	jkey := h.Itok(j)
	b := strcmp(ikey, jkey)
	switch b {
	case 1:
		ret = false
		return ret
	case 0:
		return true
	case -1:
		ret = true
		return ret
	default:
		return true
	}

}

func strcmp(a, b string) int {
	abytes := []byte(a)
	bbytes := []byte(b)
	alen := len(abytes)
	blen := len(bbytes)
	ret := byte(0)
	if alen <= blen {
		ret = 0
		for i := 0; i < alen; i++ {
			ret = abytes[i] - bbytes[i]
			if ret < 0 {
				return -1
			}
			if ret > 0 {
				return 1
			}
		}
		return -1
	} else {
		ret = 0
		for i := 0; i < blen; i++ {
			ret = abytes[i] - bbytes[i]
			if ret < 0 {
				return -1
			}
			if ret > 0 {
				return 1
			}
		}
		return 1
	}
}
