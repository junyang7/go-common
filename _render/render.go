package _render

import (
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_pb"
	"github.com/junyang7/go-common/_response"
	"net/http"
)

type Render struct {
	format   string
	response *_response.Response
}

func New() *Render {
	return &Render{
		format: "JSON",
	}
}
func (this *Render) Format(format string) *Render {
	this.format = format
	return this
}
func (this *Render) Response(response *_response.Response) *Render {
	this.response = response
	return this
}
func (this *Render) DoRpc() (o *_pb.Response) {
	o = &_pb.Response{}
	switch this.format {
	case "JSON":
		o.Response = _json.Encode(this.response)
		break
	}
	return o
}
func (this *Render) DoApi(w http.ResponseWriter) {
	switch this.format {
	case "JSON":
		w.Header().Set("content-type", "application/json")
		_, _ = w.Write(_json.Encode(this.response))
		break
	}
}
