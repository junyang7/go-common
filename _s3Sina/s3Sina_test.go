package _s3Sina

import (
	"fmt"
	"testing"
)

func TestUpload(t *testing.T) {

	{
		accessKey := ""
		secretKey := ""
		bucket := ""
		localFile := ""
		cloudFile := ""
		res := Upload(accessKey, secretKey, bucket, localFile, cloudFile, AclAuthenticatedRead, 0)
		fmt.Println(res)
	}
	{
		accessKey := ""
		secretKey := ""
		bucket := ""
		localFile := ""
		cloudFile := ""
		{
			res := Upload(accessKey, secretKey, bucket, localFile, cloudFile, AclPrivate, 60)
			fmt.Println(res)
		}
		{
			res := GetSignedUrl(accessKey, secretKey, bucket, cloudFile, 60)
			fmt.Println(res)
		}
	}

}
