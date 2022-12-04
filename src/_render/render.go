package _render

import (
	"encoding/json"
	"github.com/junyang7/go-common/src/_is"
	"github.com/junyang7/go-common/src/_response"
	"net/http"
)

type render struct {
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request) *render {
	return &render{
		w: w,
		r: r,
	}
}
func (this *render) Json(value interface{}) {
	res := _response.New()
	if _is.NotEmpty(value) {
		res.Data = value
	}
	this.w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(this.w)
	_ = e.Encode(res)
}
func (this *render) Text(value interface{}) {
	res := _response.New()
	if _is.NotEmpty(value) {
		res.Data = value
	}
	this.w.Header().Set("content-type", "application/json")
	e := json.NewEncoder(this.w)
	_ = e.Encode(res)
}
