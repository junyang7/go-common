package _server

import (
	"github.com/junyang7/go-common/src/_server/_engine"
	"github.com/junyang7/go-common/src/_server/_engine/_api"
)

func Api(conf *_api.Conf) {
	_api.Initialize(conf)
}
func Cli() {
	e := &_engine.Cli{}
	e.Serve()
}
