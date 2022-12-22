package _web

import (
	"net/http"
)

type Conf struct {
	Root string `json:"root"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type engine struct {
	conf *Conf
}

func Initialize(conf *Conf) {

	this := &engine{conf: conf}
	http.Handle("/", http.FileServer(http.Dir(conf.Root)))
	if err := http.ListenAndServe(this.conf.Ip+":"+this.conf.Port, nil); nil != err {
		panic(err)
	}

}
