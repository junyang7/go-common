package _server

import (
	"github.com/junyang7/go-common/src/_server/_engine"
	"github.com/junyang7/go-common/src/_server/_engine/_api"
	"github.com/junyang7/go-common/src/_server/_engine/_http"
	"github.com/junyang7/go-common/src/_server/_engine/_web"
)

func Api(conf *_api.Conf) {
	_api.Initialize(conf)
}
func Web(conf *_web.Conf) {
	_web.Initialize(conf)
}
func Http(conf *_http.Conf) {
	_http.Initialize(conf)
}
func Cli() {
	e := &_engine.Cli{}
	e.Serve()
}
