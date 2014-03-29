package oss

type header struct {
	// Standard HTTP Headers
	CACHE_CONTROL       string
	CONTENT_DISPOSITION string
	CONTENT_ENCODING    string
	CONTENT_LENGTH      string
	CONTENT_MD5         string
	CONTENT_TYPE        string
	DATE                string
	ETAG                string
	LAST_MODIFIED       string
	SERVER              string

	// Aliyun HTTP Headers
	ALI_PREFIX               string
	ALI_USER_METADATA_PREFIX string

	CANNED_ACL             string
	SERVER_SIDE_ENCRYPTION string
	REQUEST_ID             string
	MAGIC                  string

	COPY_SOURCE                     string
	COPY_SOURCE_IF_MATCH            string
	COPY_SOURCE_IF_NONE_MATCH       string
	COPY_SOURCE_IF_UNMODIFIED_SINCE string
	COPY_SOURCE_IF_MODIFIED_SINCE   string
	METADATA_DIRECTIVE              string

	FILE_GROUP string

	USER_METADATA_NAME   string
	USER_METADATA_AUTHOR string
}

const (
	Header = header{
		CACHE_CONTROL:       "Cache-Control",
		CONTENT_DISPOSITION: "Content-Disposition",
		CONTENT_ENCODING:    "Content-Encoding",
		CONTENT_LENGTH:      "Content-Length",
		CONTENT_MD5:         "Content-MD5",
		CONTENT_TYPE:        "Content-Type",
		DATE:                "Date",
		ETAG:                "ETag",
		LAST_MODIFIED:       "Last-Modified",
		SERVER:              "Server",

		ALI_PREFIX:               "x-oss-",
		ALI_UESR_METADATA_PREFIX: "x-oss-meta-",
		CANNED_ACL:               "x-oss-acl",
		SERVER_SIDE_ENCRYPTION:   "x-oss-server-side-encryption",
		REQUEST_ID:               "x-oss-request-id",
		MAGIC:                    "x-oss-magic",

		COPY_SOURCE:                     "x-oss-copy-source",
		COPY_SOURCE_IF_MATCH:            "x-oss-copy-source-if-match",
		COPY_SOURCE_IF_NONE_MATCH:       "x-oss-copy-source-if-none-match",
		COPY_SOURCE_IF_UNMODIFIED_SINCE: "x-oss-copy-source-if-unmodified-since",
		COPY_SOURCE_IF_MODIFIED_SINCE:   "x-oss-copy-source-if-modified-since",
		METADATA_DIRECTIVE:              "x-oss-metadata-directive",

		FILE_GROUP: "x-oss-file-group",

		USER_METADATA_NAME:   "x-oss-meta-name",
		USER_METADATA_AUTHOR: "x-oss-meta-author",
	}
)

const (
	TIME_LAYOUT = "Mon, 2 Jan 2006 15:04:05 GMT"
)
