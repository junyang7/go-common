package _mail

import (
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"gopkg.in/gomail.v2"
)

type Conf struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Passport string `json:"passport"`
	From     string `json:"from"`
}

type mail struct {
	conf    *Conf
	to      []string
	cc      []string
	subject string
	content string
	attach  []string
}

func New(conf *Conf, to []string) *mail {
	return &mail{
		conf: conf,
		to:   to,
	}
}
func (this *mail) Cc(cc []string) *mail {
	this.cc = cc
	return this
}
func (this *mail) Subject(subject string) *mail {
	this.subject = subject
	return this
}
func (this *mail) Content(content string) *mail {
	this.content = content
	return this
}
func (this *mail) Attach(attach []string) *mail {
	this.attach = attach
	return this
}
func (this *mail) Send() {
	d := gomail.NewDialer(this.conf.Host, this.conf.Port, this.conf.Username, this.conf.Passport)
	m := gomail.NewMessage()
	m.SetHeader("From", this.conf.From)
	m.SetHeader("To", this.to...)
	if len(this.cc) > 0 {
		m.SetHeader("Cc", this.cc...)
	}
	m.SetHeader("Subject", this.subject)
	m.SetBody("text/html", this.content)
	if len(this.attach) > 0 {
		for _, attach := range this.attach {
			m.Attach(attach)
		}
	}
	err := d.DialAndSend(m)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrGoMailDialerDialAndSend).
		Message(err).
		Do()
}
