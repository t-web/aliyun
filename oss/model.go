package oss

import (
	"time"
)

type Bucket struct {
	Name       string
	Owner      string
	CreateDate time.Time
}

func NewBucket(name string) *Bucket {
	ret := new(Bucket)
	ret.Name = name
	return ret
}

func (b *Bucket) String() string {
	return "OSSBucket [name=" + b.Name + ", createDate=" + b.CreateDate.String() + ", owner=" + b.Owner + "]"
}
