package _mail

import "gopkg.in/gomail.v2"

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
	if err := d.DialAndSend(m); nil != err {
		panic(err)
	}
}
