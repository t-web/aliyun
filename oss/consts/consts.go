package consts

const (
	// normal HeaderName in Metadata
	CACHE_CONTROL       = "Cache-Control"
	CONTENT_DISPOSITION = "Content-Disposition"
	CONTENT_ENCODING    = "Content-Encoding"
	CONTENT_LENGTH      = "Content-Length"
	CONTENT_MD5         = "Content-MD5"
	CONTENT_TYPE        = "Content-Type"
	DATE                = "Date"
	ETAG                = "ETag"
	SERVER              = "Server"

	HOST   = "Host"
	SCHEME = "Scheme"

	// Aliyun HeaderName in UserMetadata
	ALI_PREFIX               = "x-oss-"
	ALI_UESR_METADATA_PREFIX = "x-oss-meta-"
	CANNED_ACL               = "x-oss-acl"
	SERVER_SIDE_ENCRYPTION   = "x-oss-server-side-encryption"
	REQUEST_ID               = "x-oss-request-id"
	MAGIC                    = "x-oss-magic"
	// COPY ACTHION HEADERS
	COPY_SOURCE                     = "x-oss-copy-source"
	COPY_SOURCE_IF_MATCH            = "x-oss-copy-source-if-match"
	COPY_SOURCE_IF_NONE_MATCH       = "x-oss-copy-source-if-none-match"
	COPY_SOURCE_IF_UNMODIFIED_SINCE = "x-oss-copy-source-if-unmodified-since"
	COPY_SOURCE_IF_MODIFIED_SINCE   = "x-oss-copy-source-if-modified-since"
	METADATA_DIRECTIVE              = "x-oss-metadata-directive"

	FILE_GROUP = "x-oss-file-group"

	USER_METADATA_NAME   = "x-oss-meta-name"
	USER_METADATA_AUTHOR = "x-oss-meta-author"

	USER_AUTHORIZATION = "Authorization"
)

const (
	RFC2616_AUTHORIZATION       = "Authorization"
	RFC2616_CACHE_CONTROL       = "Cache-Control"
	RFC2616_CONTENT_DISPOSITION = "Content-Disposition"
	RFC2616_CONTENT_ENCODING    = "Content-Encoding"
	RFC2616_CONTENT_LENGTH      = "Content-Length"
	RFC2616_CONTENT_MD5         = "Content-MD5"
	RFC2616_CONTENT_TYPE        = "Content-Type"
	RFC2616_CONNECTION          = "Connection"
	RFC2616_DATE                = "Date"
	RFC2616_HOST                = "Host"
	RFC2616_LAST_MODIFIED       = "Last-Modified"
	RFC2616_METHOD              = "Method"
	RFC2616_USER_ANGENT         = "User-Agent"
)

const (
	RFC1123G = "Mon, 02 Jan 2006 15:04:05 GMT"
)

const (
	DOMAIN_DEFAULT           = "oss.aliyuncs.com"
	DOMAIN_HANGZHOU          = "oss-cn-hangzhou.aliyuncs.com"
	DOMAIN_QINGDAO           = "oss-cn-qingdao.aliyuncs.com"
	DOMAIN_HANGZHOU_INTERNAL = "oss-cn-hangzhou-internal.aliyuncs.com"
	DOMAIN_QINGDAO_INTERNAL  = "oss-cn-qingdao-internal.aliyuncs.com"
)
