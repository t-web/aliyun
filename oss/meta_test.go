package oss_test

import (
	"github.com/stduolc/aliyun/oss"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"testing"
)

func testT(t *testing.T) {
	heads := make(oss.UserMetadata)
	for k, v := range heads {
		heads[i].SetKey(strconv.FormatInt(int64(i*3+10), 13))
	}
	sort.Sort(heads)
	for _, v := range heads {
		t.Log(v)
	}
}

func testTT(t *testing.T) {
	req := http.Request{}
	req.Method = "PUT"
	req.URL = &url.URL{Path: "/nelson"}
	req.Header = make(map[string][]string)
	req.Proto = "HTTP/1.0"
	req.Header.Set("Host", "oss-example.oss-cn-hangzhou.aliyuncs.com")
	req.Header.Set("Content-Md5", "c8fdb181845a4ca6b8fec737b3581d76")
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("Date", "Thu, 17 Nov 2005 18:49:58 GMT")
	req.Header.Set("X-OSS-Meta-Author", "foo@bar.com")
	req.Header.Set("X-OSS-Magic", "abracadabra")

	meta := oss.Metadata{}
	meta.SetAuth(&oss.Auth{
		AccessKey: "44CF9590006BF252F707",
		SecretKey: "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV",
	})
	meta.SetRequest(&req)
	t.Log("Request: ", req)
	for _, v := range meta.GetUserMetadata() {
		t.Log("Meta: ", *v)
	}
}
