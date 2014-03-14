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
	req.URL = url.URL{"Path": "/nelson"}
	req.Proto = "HTTP/1.0"
	req.URL.Host = "oss-example.oss-cn-hangzhou.aliyuncs.com"
	req.Header.Set("Host", "oss-example.oss-cn-hangzhou.aliyuncs.com")
	req.Header.Set("Content-Md5", "c8fdb181845a4ca6b8fec737b3581d76")
	req.Header.Set("Content-Type", "text/html")
	req.Header.Set("Date", "Thu, 17 Nov 2005 18:49:58 GMT")
	req.Header.Set("X-OSS-Meta-Author", "foo@bar.com")
	req.Header.Set("X-OSS-Magic", "abracadabra")
}
