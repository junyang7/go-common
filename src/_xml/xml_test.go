package _xml

import (
	"encoding/xml"
	"fmt"
	"github.com/junyang7/go-common/src/_string"
	"testing"
)

func TestDecode(t *testing.T) {
	{

		txt := `
<?xml="" version="1.0" encoding="UTF-8" ?=""> 
<user>
 	<info>
  		<username>225531</username>
 		<email>junyang7</email>
 		<uid>7612A97A-FA9C-4C5A-A903-F914E19B0E4F</uid>
 		<fullemail>junyang7@staff.weibo.com</fullemail>
 		<name>郭俊阳</name>
 		<erpname>郭俊阳225531</erpname>
 		<organization>电商闭环技术部</organization>
 		<organizationt3>新浪集团_微博_微博增值业务部</organizationt3>
 		<companycode>SINA</companycode>
 		<companycodes></companycodes>
 		<telephone>17772120175</telephone>
	</info>
  <acl></acl>
</user>
`
		txt = _string.ReplaceAll(txt, `xml=""`, `xml`)
		txt = _string.ReplaceAll(txt, `?=""`, `?`)
		type User struct {
			XMLName xml.Name `xml:"user"`
			Info    struct {
				Username       string `xml:"username"`
				Email          string `xml:"email"`
				UID            string `xml:"uid"`
				FullEmail      string `xml:"fullemail"`
				Name           string `xml:"name"`
				ERPName        string `xml:"erpname"`
				Organization   string `xml:"organization"`
				OrganizationT3 string `xml:"organizationt3"`
				CompanyCode    string `xml:"companycode"`
				CompanyCodes   string `xml:"companycodes"`
				Telephone      string `xml:"telephone"`
			} `xml:"info"`
			ACL string `xml:"acl"`
		}

		var user User
		Decode([]byte(txt), &user)

		fmt.Println("username", user.Info.Username)
		fmt.Println("email", user.Info.Email)
		fmt.Println("fullemail", user.Info.FullEmail)
		fmt.Println("name", user.Info.Name)

	}
}
