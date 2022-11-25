package _server

import (
	"github.com/junyang7/go-common/src/_context"
	"github.com/junyang7/go-common/src/_response"
	"github.com/junyang7/go-common/src/_server/_conf"
	"net/http"
)

func Api() {
	e := &server{mode: "api"}
	e.Serve()
}
func Cli() {
	e := &server{mode: "cli"}
	e.Serve()
}

type server struct {
	mode string
	conf _conf.Conf
}

func (this *server) Serve() {
	this.conf = _conf.Conf{}
	this.conf.Ip = "0.0.0.0"
	this.conf.Port = "8888"
	if err := http.ListenAndServe(":8888", this); nil != err {
		panic(err)
	}
}
func (this *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e := &engine{
		response: &_response.Response{},
		ctx:      &_context.Context{W: w, R: r},
	}
	e.do()
}
