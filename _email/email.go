package _email

import (
	"github.com/junyang7/go-common/_interceptor"
	"gopkg.in/gomail.v2"
)

type mail struct {
	host       string
	port       int
	username   string
	password   string
	from       string
	toList     []string
	subject    string
	content    string
	ccList     []string
	attachList []string
}

func New(host string, port int, username string, password string, from string, toList ...string) *mail {
	return &mail{
		host:       host,
		port:       port,
		username:   username,
		password:   password,
		from:       from,
		toList:     toList,
		subject:    "",
		content:    "",
		ccList:     []string{},
		attachList: []string{},
	}
}
func (this *mail) Subject(subject string) *mail {
	this.subject = subject
	return this
}
func (this *mail) Content(content string) *mail {
	this.content = content
	return this
}
func (this *mail) Cc(ccList ...string) *mail {
	this.ccList = append(this.ccList, ccList...)
	return this
}
func (this *mail) Attach(attachList ...string) *mail {
	this.attachList = append(this.attachList, attachList...)
	return this
}
func (this *mail) Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", this.from)
	m.SetHeader("To", this.toList...)
	m.SetHeader("Subject", this.subject)
	m.SetBody("text/html", this.content)
	if len(this.ccList) > 0 {
		m.SetHeader("Cc", this.ccList...)
	}
	if len(this.attachList) > 0 {
		for _, attach := range this.attachList {
			m.Attach(attach)
		}
	}
	d := gomail.NewDialer(this.host, this.port, this.username, this.password)
	if err := d.DialAndSend(m); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
}
