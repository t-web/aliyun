package oss

import (
	"strings"
)

var mimeTypes = map[string]string{
	".css":  "text/css; charset=utf-8",
	".gif":  "image/gif",
	".htm":  "text/html; charset=utf-8",
	".html": "text/html; charset=utf-8",
	".jpg":  "image/jpeg",
	".js":   "application/x-javascript",
	".pdf":  "application/pdf",
	".png":  "image/png",
	".xml":  "text/xml; charset=utf-8",
	".mp4":  "video/mp4",
}

func TypeByExtension(ext string) string {
	t, ok := mimeTypes[ext]
	if ok == nil {
		return t
	} else {
		return "unkown"
	}
}

func TypeByFileName(filename string) string {
	i := strings.LastIndex(filename, ".")
	sub := filename[i:]
	return TypeByExtension(sub)
}
