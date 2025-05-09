package _context

import (
	"encoding/json"
	"fmt"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_parameter"
	"github.com/junyang7/go-common/_render"
	"github.com/junyang7/go-common/_response"
	"github.com/junyang7/go-common/_unixMilli"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

type Context struct {
	TimeS   int64
	w       http.ResponseWriter
	r       *http.Request
	debug   bool
	STORE   map[string]interface{}
	GET     map[string]string
	POST    map[string]string
	REQUEST map[string]string
	COOKIE  map[string]*http.Cookie
	SERVER  map[string]string
	ENV     map[string]string
	BODY    []byte
	FILE    map[string][]*multipart.FileHeader
}

func New(w http.ResponseWriter, r *http.Request, debug bool) *Context {
	this := &Context{
		TimeS:   _unixMilli.Get(),
		w:       w,
		r:       r,
		debug:   debug,
		STORE:   map[string]interface{}{},
		GET:     map[string]string{},
		POST:    map[string]string{},
		REQUEST: map[string]string{},
		COOKIE:  map[string]*http.Cookie{},
		SERVER:  map[string]string{},
		ENV:     map[string]string{},
		BODY:    []byte{},
		FILE:    map[string][]*multipart.FileHeader{},
	}
	for k, v := range this.r.Header {
		this.SERVER[strings.ToLower(k)] = v[0]
	}
	this.SERVER["method"] = this.r.Method
	this.SERVER["path"] = this.r.URL.Path
	this.SERVER["host"] = this.r.Host
	this.SERVER["protocol"] = this.r.Proto
	this.SERVER["referer"] = this.r.Referer()
	this.SERVER["user-agent"] = this.r.UserAgent()
	for _, v := range this.r.Cookies() {
		this.COOKIE[v.Name] = v
	}
	for k, v := range this.r.URL.Query() {
		this.GET[k] = v[0]
	}
	contentType := ""
	contentTypePartList := strings.Split(strings.ToLower(this.r.Header.Get("content-type")), ";")
	if len(contentTypePartList) > 0 {
		contentType = contentTypePartList[0]
	}
	switch contentType {
	case "application/x-www-form-urlencoded":
		if err := this.r.ParseForm(); nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		for k, v := range this.r.PostForm {
			this.POST[k] = v[0]
		}
		break
	case "multipart/form-data":
		if err := this.r.ParseMultipartForm(32 << 20); nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		for k, v := range this.r.MultipartForm.Value {
			this.POST[k] = v[0]
		}
		for k, v := range this.r.MultipartForm.File {
			this.FILE[k] = v
		}
		break
	case "application/json":
		b, err := io.ReadAll(this.r.Body)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		defer func(ctx *Context) {
			if err := ctx.r.Body.Close(); nil != err {
				_interceptor.Insure(false).Message(err).Do()
			}
		}(this)
		this.BODY = b
		var post map[string]interface{}
		if err := json.Unmarshal(this.BODY, &post); nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		for k, v := range post {
			this.POST[k] = fmt.Sprintf("%v", v)
		}
		break
	default:
		break
	}
	for k, v := range this.GET {
		this.REQUEST[k] = v
	}
	for k, v := range this.POST {
		this.REQUEST[k] = v
	}
	return this
}

func (this *Context) GetValue(k string, v string) string {
	if v, ok := this.GET[k]; ok {
		return v
	}
	return v
}
func (this *Context) GetValueAll() map[string]string {
	return this.GET
}
func (this *Context) GetParameter(k string) *_parameter.Parameter {
	if _, ok := this.GET[k]; ok {
		return _parameter.New(k, this.GET[k])
	}
	return _parameter.New(k, nil)
}
func (this *Context) PostValue(k string, v string) string {
	if v, ok := this.POST[k]; ok {
		return v
	}
	return v
}
func (this *Context) PostValueAll() map[string]string {
	return this.POST
}
func (this *Context) PostParameter(k string) *_parameter.Parameter {
	if _, ok := this.POST[k]; ok {
		return _parameter.New(k, this.POST[k])
	}
	return _parameter.New(k, nil)
}
func (this *Context) RequestValue(k string, v string) string {
	if v, ok := this.REQUEST[k]; ok {
		return v
	}
	return v
}
func (this *Context) RequestValueAll() map[string]string {
	return this.REQUEST
}
func (this *Context) RequestParameter(k string) *_parameter.Parameter {
	if _, ok := this.REQUEST[k]; ok {
		return _parameter.New(k, this.REQUEST[k])
	}
	return _parameter.New(k, nil)
}
func (this *Context) Cookie(k string) *http.Cookie {
	if v, ok := this.COOKIE[k]; ok {
		return v
	}
	return nil
}
func (this *Context) CookieValue(k string, v string) string {
	if v, ok := this.COOKIE[k]; ok {
		return v.Value
	}
	return v
}
func (this *Context) CookieValueAll() map[string]string {
	v := map[string]string{}
	for _, c := range this.COOKIE {
		v[c.Name] = c.Value
	}
	return v
}
func (this *Context) CookieParameter(k string) *_parameter.Parameter {
	if _, ok := this.COOKIE[k]; ok {
		return _parameter.New(k, this.COOKIE[k].Value)
	}
	return _parameter.New(k, nil)
}
func (this *Context) ServerValue(k string, v string) string {
	if v, ok := this.SERVER[k]; ok {
		return v
	}
	return v
}
func (this *Context) ServerValueAll() map[string]string {
	return this.SERVER
}
func (this *Context) ServerParameter(k string) *_parameter.Parameter {
	if _, ok := this.SERVER[k]; ok {
		return _parameter.New(k, this.SERVER[k])
	}
	return _parameter.New(k, nil)
}
func (this *Context) File(k string) []*multipart.FileHeader {
	if f, ok := this.FILE[k]; ok && len(f) > 0 {
		return f
	}
	return nil
}
func (this *Context) Body() []byte {
	return this.BODY
}
func (this *Context) Bind(v interface{}) {
	if len(this.BODY) > 0 {
		_json.Decode(this.BODY, v)
	}
}
func (this *Context) SetHeader(k string, v string) *Context {
	this.w.Header().Set(k, v)
	return this
}
func (this *Context) SetCookie(cookie *http.Cookie) *Context {
	http.SetCookie(this.w, cookie)
	return this
}
func (this *Context) JSON(data any) {
	res := _response.New()
	if nil != data {
		switch data.(type) {
		case *_response.Response:
			res = data.(*_response.Response)
		default:
			res.Data = data
		}
	}
	if !this.debug {
		res.File = ""
		res.Line = 0
	}
	res.Time = _unixMilli.Get()
	res.Consume = res.Time - this.TimeS
	_render.JSON(this.w, res)
}
func (this *Context) REDIRECT(uri string) {
	this.w.Header().Set("Location", uri)
	this.w.WriteHeader(301)
}
