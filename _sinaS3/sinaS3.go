package _sinaS3

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_sinaS3/sdk"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const (
	AclPrivate           = "private"
	AclPublicRead        = "public-read"
	AclPublicReadWrite   = "public-read-write"
	AclAuthenticatedRead = "authenticated-read"
)

func Upload(accessKey string, secretKey string, bucket string, localFile string, cloudFile string, acl string, ttl int64) string {
	b := sdk.NewSCS(accessKey, secretKey, "https://intra-d.sinastorage.com").Bucket(bucket)
	sdkAcl := sdk.PublicReadWrite
	switch acl {
	case AclPrivate:
		sdkAcl = sdk.Private
		break
	case AclPublicRead:
		sdkAcl = sdk.PublicRead
		break
	case AclPublicReadWrite:
		sdkAcl = sdk.PublicReadWrite
		break
	case AclAuthenticatedRead:
		sdkAcl = sdk.AuthenticatedRead
		break
	default:
		_interceptor.
			Insure(false).
			Message("参数错误：acl").
			Do()
	}
	err := b.Put(cloudFile, localFile, sdkAcl)
	_interceptor.
		Insure(nil == err).
		Message(err).
		Do()
	if AclPrivate == acl {
		return GetSignedUrl(accessKey, secretKey, bucket, cloudFile, ttl)
	}
	return "https://" + bucket + "/" + cloudFile

}
func GetSignedUrl(accessKey string, secretKey string, bucket string, cloudFile string, ttl int64) string {
	b := sdk.NewSCS(accessKey, secretKey, "https://gslb.sinastorage.cn").Bucket(bucket)
	return b.SignURL(cloudFile, time.Now().Add(time.Second*time.Duration(ttl)))
}
func UploadByCtxFile(accessKey string, secretKey string, bucket string, fileHeader *multipart.FileHeader, cloudFile string, acl string, ttl int64) string {
	src, err := fileHeader.Open()
	_interceptor.
		Insure(nil == err).
		Message(err).
		Do()
	defer src.Close()
	tmp, err := os.CreateTemp("", "s3Sina-*"+filepath.Ext(fileHeader.Filename))
	_interceptor.
		Insure(nil == err).
		Message(err).
		Do()
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)
	_, err = io.Copy(tmp, src)
	_interceptor.
		Insure(nil == err).
		Message(err).
		Do()
	err = tmp.Close()
	_interceptor.
		Insure(nil == err).
		Message(err).
		Do()
	return Upload(accessKey, secretKey, bucket, tmpPath, cloudFile, acl, ttl)
}
