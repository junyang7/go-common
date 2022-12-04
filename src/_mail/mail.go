package _mail

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"gopkg.in/gomail.v2"
)

type mail struct {
	host     string
	port     int
	username string
	password string
	from     string
	to       []string
	cc       []string
	subject  string
	content  string
	attach   []string
}

func New(host string, port int, username string, password string, from string, to []string, subject string, content string) *mail {
	return &mail{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
		to:       to,
		subject:  subject,
		content:  content,
	}
}
func (this *mail) Cc(cc ...string) *mail {
	this.cc = cc
	return this
}
func (this *mail) Attach(attach ...string) *mail {
	this.attach = attach
	return this
}
func (this *mail) Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", this.from)
	m.SetHeader("To", this.to...)
	m.SetHeader("Cc", this.cc...)
	m.SetHeader("Subject", this.subject)
	m.SetBody("text/html", this.content)
	for _, attach := range this.attach {
		m.Attach(attach)
	}
	d := gomail.NewDialer(this.host, this.port, this.username, this.password)
	err := d.DialAndSend(m)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrGoMailDialerDialAndSend).
		Message(err).
		Do()
}
