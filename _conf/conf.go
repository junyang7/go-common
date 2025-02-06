package _conf

import "github.com/junyang7/go-common/_sql"

type conf struct {
	Env   string
	Debug bool
	Http  struct {
		Addr   string
		Root   string
		Origin []string
	}
	Sql   map[string]*_sql.Business
	Email struct {
		Host     string
		Port     int
		Username string
		Password string
		From     string
	}
}

var Conf = &conf{}
