// sdkTest project main.go
// @author:shixi_mingzhe1 2018/04/17
package sdk

import "fmt"
import "os"
import "time"

//bytes����תString
func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

/*
	���Բ����������̲��ԣ�����bucket->bucket�йز���->�ϴ�object->object�йز���->��Ƭ�ϴ�->ɾ��object->ɾ��bucket��
	�ڶԸո��ϴ���object����ʱ���ܻ����no such key���⣬�����ӳ���ɵģ���Ӱ�칦��
*/
func main() {

	//���������AccessKey��SecretKey
	var AccessKey string = ""
	var SecretKey string = ""
	/*
		golang ��Ҫ��URL�����Э�����ƣ�����http����https������ϡ�
	*/
	var DEFAULT_DOMAIN string = "https://intra-d.sinastorage.com"
	var BucketName string = "adonis-3"
	var LocalFile string = "C:/Users/shixi_mingzhe1/Desktop/daily_life.txt"
	var CloudFile string = "GoTest/daily_life.txt"
	//	var sha1 string = "370c93e3f87ad26e31db080256914ea3ae2d25d1"
	//	var MultipartLocalFile string = "D:/CentOSMirror/CentOS-7-x86_64-DVD-1708.iso"
	var MultipartLocalFile string = "D:/paper/yaoyan.zip"
	var MultipartCloudFile string = "multipart/multipart.iso"
	//init scs
	scs := NewSCS(AccessKey, SecretKey, DEFAULT_DOMAIN)
	fmt.Println("Create scs---", scs)
	// bucket
	// create bucket&&set acl
	bucket := scs.Bucket(BucketName)
	put_bucket_err := bucket.PutBucket(Private)
	fmt.Println("Create bucket---", BucketName, put_bucket_err)
	//set bucket acl
	acl_bucket := map[string][]string{"GRPS000000ANONYMOUSE": []string{"read", "read_acp", "write_acp"}}
	set_bucket_acl_err := bucket.SetBucketAcl(acl_bucket)
	fmt.Println("set bucket acl", set_bucket_acl_err)
	// list bucket(service)
	list_bucket, list_bucket_err := bucket.ListBucket()
	fmt.Println("list bucket---", byteString(list_bucket), list_bucket_err)
	// list  Object
	list, list_object_err := bucket.ListObject("", "", "", 10)
	fmt.Println("List objects---", BucketName, byteString(list), list_object_err)
	// get bucket acl info
	acl_data, get_acl_err := bucket.GetBucketInfo("acl")
	fmt.Println("Acl Info---", BucketName, byteString(acl_data), get_acl_err)
	// get	bucket meta info
	meta_data, get_meta_err := bucket.GetBucketInfo("meta")
	fmt.Println("Meta Info---", BucketName, byteString(meta_data), get_meta_err)
	// delete  bucket
	//	del_bucket_err := bucket.DelBucket()
	//	fmt.Println("Delete bucket---", BucketName, del_bucket_err)
	// object
	// put object&set ACL
	put_object_err := bucket.Put(CloudFile, LocalFile, PublicRead)
	fmt.Println("Put object---", CloudFile, put_object_err)
	//copy object
	copy_err := bucket.Copy("Copy/copy.txt", "adonis-1", "adonis-1test.txt")
	fmt.Println("copy object---Copy/copy.txt", copy_err)
	// put object relax
	relax_err := bucket.Relax("Relax/relax.txt", LocalFile, PublicReadWrite)
	fmt.Println("relax object---Relax/relax.txt", relax_err)
	// download object contents
	bucket_1 := scs.Bucket("adonis-1")
	object_content, con_err := bucket_1.Get("adonis-1test.txt")
	fp, err := os.Create("C:/Users/shixi_mingzhe1/Desktop/go_download_test.txt")
	download_result, download_err := fp.WriteString(byteString(object_content))
	fp.Close()
	fmt.Println("object contents---", byteString(object_content), con_err, download_result, download_err)
	// delete object
	del_err := bucket.Del("Relax/relax.txt")
	fmt.Println("delete object---Relax/relax.txt", del_err)
	//multipart upload
	fmt.Println("init multipart starting!")
	multi, err := bucket.InitMulti(MultipartCloudFile)
	if err != nil {
		return
	}
	fmt.Println("upload multipart starting!")
	part_info, err := multi.PutPart(MultipartLocalFile, Private, 1024*1024*5)
	if err != nil {
		return
	}
	listpart, err := multi.ListPart()
	fmt.Println("multiparts---------", listpart)
	if err != nil {
		return
	}
	for k, v := range listpart {
		if part_info[k].ETag != v.ETag {
			fmt.Println("multipart not match!")
		}
	}
	com_err := multi.Complete(listpart)
	if com_err == nil {
		fmt.Println("multipart upload successfully!")
	}
	// put object meta��set object acl
	meta := map[string]string{"x-amz-meta-adonis": "adonis", "x-amz-meta-test": "20180417"}
	put_meta_err := bucket_1.PutMeta("adonis-1test.txt", meta)
	acl := map[string][]string{"GRPS000000ANONYMOUSE": []string{"read", "read_acp", "write", "write_acp"}}
	put_acl_err := bucket_1.PutAcl("adonis-1test.txt", acl)
	fmt.Println("put object Meta and Acl", "adonis-1test.txt", put_meta_err, put_acl_err)
	// get object Acl&Meta info
	object_metainfo, object_metainfo_err := bucket_1.GetInfo("adonis-1test.txt", "meta")
	object_aclinfo, object_aclinfo_err := bucket_1.GetInfo("adonis-1test.txt", "acl")
	fmt.Println("Meta and Acl---", "adonis-1test.txt", byteString(object_metainfo), byteString(object_aclinfo), object_aclinfo_err, object_metainfo_err)

	// Authed Url with Expires
	authed_uri := bucket.SignURL(MultipartCloudFile, time.Now().Add(15*time.Second))
	// Url
	uri := bucket.URL(MultipartCloudFile)
	fmt.Println("Url for:", MultipartCloudFile, " Authed:", authed_uri, " Unauthed:", uri)

	del_err_0 := bucket.Del(MultipartCloudFile)
	if del_err_0 != nil {
		fmt.Println("del failed")
	}
	del_err_1 := bucket.Del(CloudFile)
	if del_err_1 != nil {
		fmt.Println("del failed")
	}
	del_err_2 := bucket.Del("Copy/copy.txt")
	if del_err_2 != nil {
		fmt.Println("del failed")
	}
	//	del_err_3 := bucket.DelBucket()
	//	if del_err_3 != nil {
	//		fmt.Println("del failed")
	//	}

}
