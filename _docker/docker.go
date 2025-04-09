package _docker

import (
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_cmd"
	"github.com/junyang7/go-common/_json"
	"strings"
)

type Host struct {
	Port string `json:"HostPort"`
}

func InspectPort(name string) (o map[int64]int64) {
	var t map[string][]*Host
	b := _cmd.Execute(`docker`, `inspect`, `-f`, `{{json .NetworkSettings.Ports}}`, name)
	_json.Decode(b, &t)
	o = map[int64]int64{}
	for k, v := range t {
		o[_as.Int64(strings.Split(k, "/")[0])] = _as.Int64(v[0].Port)
	}
	return o
}
