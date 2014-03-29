package oss

import (
	"io"
	"os"
)

type PutObjectRequest struct {
	// private members
	BucketName              string
	Key                     string
	File                    os.File
	InputStream             io.Reader
	Metadata                ObjectMetadata
	CannedAcl               CannedAccessControlList
	AccessControlList       AccessControlList
	StorageClass            string
	GeneralProgressListener ProgressListener
	RedirectLocation        string
}

func NewPutObjectRequest(args [...]interface{}) (ret *PutObjectRequest, err error) {
	n := len(args)
	if n == 3 {
		// type(args[2]) ==
	}
	switch n {
	case 3:
		switch args[2].(type) {
		case os.File:
			ret := constructorA(args[0], args[1], args[2])
			return ret, nil
		case string:
			ret := constructorB(args[0], args[1], args[2])
			return ret, nil
		default:
			return nil, Error("ErrorNum 2. Plz chech the document for more info!")
		}
	case 4:
	default:
		return nil, Error("ErrorNum 1. Plz check the document for more info!")
	}
	return nil, Error("ErrorNum 3. Plz check the document for more info!")
}

func (r *PutObjectRequest) WithBucketName(bucketName string) *PutObjectRequest {
	if r == nil {
		return nil
	}
	r.BucketName = bucketName
	return r
}

func (r *PutObjectRequest) WithKey(key string) *PutObjectRequest {
	if r == nil {
		return nil
	}
	r.Key = key
	return r
}

func (r *PutObjectRequest) WithStorageClass(storageClass string) *PutObjectRequest {
	if r == nil {
		return nil
	}
	r.StorageClass = storageClass
	return r
}

func (r *PutObjectRequest) WithStorageClass(storageClass StorageClass) *PutObjectRequest {
	if r == nil {
		return nil
	}
	r.StorageClass = storageClass.String()
	return r
}

func constructorA(bucketName, key string, file os.File) (ret *PutObjectRequest) {
	ret = new(PutObjectRequest)
	ret.BucketName = bucketName
	ret.Key = key
	ret.File = file
	return ret
}

//TODO
func constructorB(bucketName, key, redirectLocation string) (ret *PutObjectRequest) {
	ret = new(PutObjectRequest)
	ret.BucketName = bucketName
	ret.Key = key
	ret.RedirectLocation = redirectLocation
	return ret
}

func constructorC(bucketName, key string, inputStream io.Reader, metadata ObjectMetadata) (ret *PutObjectRequest) {
	ret = new(PutObjectRequest)
	ret.BucketName = bucketName
	ret.Key = key
	ret.InputStream = inputStream
	ret.Metadata = metadata
	return ret
}
