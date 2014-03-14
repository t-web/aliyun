package oss

import (
	"net/http"
	"sort"
	"strings"
)

type Head struct {
	Key   string
	Value string
}

type Heads []*Head

type CanonicalizedOSSHeaders struct {
	Heads
}

type CanonicalizedResource struct {
}

type Meta struct {
	VERB        string
	ContentMD5  string
	ContentType string
	Date        string
	OSSHeaders  CanonicalizedOSSHeaders
	Resource    CanonicalizedResource

	Signature string
}

func (c *CanonicalizedOSSHeaders) SetOSSHeaders(req *http.Request) {
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
	c.Heads = heads
}

func (heads Heads) Len() int {
	return len(heads)
}

func (heads Heads) Swap(i, j int) {
	heads[i], heads[j] = heads[j], heads[i]
}

func (h Heads) Less(i, j int) bool {
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
