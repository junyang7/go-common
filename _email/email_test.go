package _email

import "testing"

func TestMail_Send(t *testing.T) {
	//{
	//	New(
	//		"smtp.126.com",
	//		25,
	//		"guoguo_1234567890@126.com",
	//		"FZQJRCOHGWXSESMU",
	//		"guoguo_1234567890@126.com",
	//		"507979696@qq.com").Send()
	//}
	//{
	//	New(
	//		"smtp.126.com",
	//		25,
	//		"guoguo_1234567890@126.com",
	//		"FZQJRCOHGWXSESMU",
	//		"guoguo_1234567890@126.com",
	//		"507979696@qq.com").
	//		Subject("subject").
	//		Content("content").
	//		Cc("507979696@qq.com").
	//		Attach("/Users/junyang7/ziji/go-common/go.mod").
	//		Send()
	//}
	//{
	//	New(
	//		"smtp.126.com",
	//		25,
	//		"guoguo_1234567890@126.com",
	//		"FZQJRCOHGWXSESMU",
	//		"guoguo_1234567890@126.com",
	//		"507979696@qq.com").
	//		Subject("subject").
	//		Content("<a href='https://www.baidu.com'>跳转</a>").
	//		Send()
	//}
	{
		New(
			"smtp.126.com",
			25,
			"guoguo_1234567890@126.com",
			"FZQJRCOHGWXSESMU",
			"guoguo_1234567890@126.com",
			"507979696@qq.com").
			Subject("subject").
			Content("<img src='https://www.baidu.com/img/flexible/logo/pc/result@2.png'/>").
			Send()
	}
}
