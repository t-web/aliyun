package oss

import (
// "time"
)

type Bucket struct {
	Name        string
	Owner       string
	Acl         string
	Prefix      string
	Marker      string
	MaxKeys     string
	Delimiter   string
	IsTruncated bool

	Contents []Content
}

type Content struct {
	Key          string
	LastModified string
	ETag         string
	Type         string
	Size         string
	StorageClass string
}

type Owner struct {
	Id          string
	DisplayName string
}

func NewBucket(name string) *Bucket {
	ret := new(Bucket)
	ret.Name = name
	return ret
}

func (b *Bucket) String() string {
	return "OSSBucket [name=" + b.Name + ", owner=" + b.Owner + ", acl=" + b.Acl + "]"
}
