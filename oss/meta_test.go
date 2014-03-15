package oss_test

import (
	"github.com/stduolc/aliyun/oss"
	"net/http"
	"net/url"
	"sort"
	"testing"
)

func TestT(t *testing.T) {
	heads := oss.Heads{&oss.Head{"ba", "1"}, &oss.Head{"b0", "2"}, &oss.Head{"a", "3"}}
	sort.Sort(heads)
	for _, v := range heads {
		t.Log(v)
	}
}

func TestTT(t *testing.T) {
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

	meta := oss.Meta{}
	meta.SetAuth(&oss.Auth{
		AccessKey: "44CF9590006BF252F707",
		SecretKey: "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV",
	})
	meta.SetRequest(&req)
	t.Log("Request: ", req)
	t.Log("Meta: ", meta.Resource)
}
