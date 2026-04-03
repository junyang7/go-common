package _s3Sina

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_s3Sina/sdk"
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
