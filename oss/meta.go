package oss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type Auth struct {
	AccessKey string
	SecretKey string
}

type Head struct {
	Key   string
	Value string
}

type Heads []*Head

type Meta struct {
	Auth        *Auth
	Method      string
	ContentMD5  string
	ContentType string
	Date        string
	OSSHeaders  Heads
	Resource    string

	Signature string
}

func (m *Meta) SetAuth(auth *Auth) {
	m.Auth = auth
}

func (m *Meta) SetRequest(req *http.Request) {
	m.setMethod(req)
	m.setContentMD5(req)
	m.setContentType(req)
	m.setDate(req)
	m.setOSSHeaders(req)
	m.setResource(req)
	m.signr()
}

func (m *Meta) setMethod(req *http.Request) {
	m.Method = req.Method
}

func (m *Meta) setContentMD5(req *http.Request) {
	m.ContentMD5 = req.Header.Get("content-md5")
}

func (m *Meta) setContentType(req *http.Request) {
	m.ContentType = req.Header.Get("content-type")
}

func (m *Meta) setDate(req *http.Request) {
	m.Date = req.Header.Get("Date")
}

func (m *Meta) setOSSHeaders(req *http.Request) {
	heads := make(Heads, 0)
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
			heads = append(heads, &Head{key, val})
		}
	}
	sort.Sort(heads)
	m.OSSHeaders = heads
}

func (m *Meta) setResource(req *http.Request) {
	tmp := req.Header.Get("Host")
	sp := strings.Split(tmp, ".")
	bucket := sp[0]
	name := req.URL.Path
	name = fmt.Sprint("/", bucket, name)
	m.Resource = name
}

func (m *Meta) signr() {
	str := fmt.Sprint(m.Method, "\n", m.ContentMD5, "\n", m.ContentType, "\n", m.Date, "\n")
	for _, v := range m.OSSHeaders {
		str = fmt.Sprint(str, v.Key, ":", v.Value, "\n")
	}
	str = fmt.Sprint(str, m.Resource)
	fmt.Println(str)
	mac := hmac.New(sha1.New, []byte(m.Auth.SecretKey))
	mac.Write([]byte(str))
	dst := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	m.Signature = dst
	fmt.Println(dst)
}

func (heads Heads) Len() int {
	return len(heads)
}

func (heads Heads) Swap(i, j int) {
	heads[i], heads[j] = heads[j], heads[i]
}

func (h Heads) Less(j, i int) bool {
	ret := true
	ikey := h[i].Key
	jkey := h[j].Key
	ilen := len(ikey)
	jlen := len(jkey)
	if ilen <= jlen {
		for i := 0; i < ilen; i++ {
			ret = (ikey[i] < jkey[i])
			if ret == false {
				return ret
			}
		}
		return ret
	} else {
		ret = false
		for i := 0; i < jlen; i++ {
			ret = (ikey[i] > jkey[i])
			if ret == true {
				return ret
			}
		}
		return ret
	}
}
