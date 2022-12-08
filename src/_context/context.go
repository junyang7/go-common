package _context

import (
	"fmt"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_is"
	"github.com/junyang7/go-common/src/_parameter"
	"github.com/junyang7/go-common/src/_render"
	"github.com/junyang7/go-common/src/_response"
	"mime/multipart"
	"net/http"
	"strings"
)

type Context struct {
	w       http.ResponseWriter
	r       *http.Request
	get     map[string]string
	post    map[string]string
	request map[string]string
	cookie  map[string]string
	server  map[string]string
}

func New(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		w: w,
		r: r,
	}
}

func (this *Context) prepareGet() {
	if nil != this.get {
		return
	}
	this.get = map[string]string{}
	for k, v := range this.r.URL.Query() {
		this.get[k] = v[0]
	}
}
func (this *Context) preparePost() {
	if nil != this.post {
		return
	}
	this.post = map[string]string{}
	err := this.r.ParseForm()
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHttpRequestParseForm).
		Message(err).
		Do()
	contentType := strings.ToLower(this.r.Header.Get("content-type"))
	if -1 != strings.Index(contentType, "application/x-www-form-urlencoded") {
		for k, v := range this.r.PostForm {
			this.post[k] = v[0]
		}
	}
	if -1 != strings.Index(contentType, "multipart/form-data") {
		err := this.r.ParseMultipartForm(32 << 20)
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrHttpRequestParseMultipartForm).
			Message(err).
			Do()
		for k, v := range this.r.PostForm {
			this.post[k] = v[0]
		}
	}
	if -1 != strings.Index(contentType, "application/json") {
		post := map[string]interface{}{}
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrJsonNewDecoderDecode).
			Message(err).
			Do()
		err := this.r.Body.Close()
		_interceptor.Insure(nil == err).
			CodeMessage(_codeMessage.ErrHttpRequestBodyClose).
			Message(err).
			Do()
		for k, v := range post {
			this.post[k] = fmt.Sprintf("%v", v)
		}
	}
}
func (this *Context) prepareRequest() {
	if nil != this.request {
		return
	}
	this.request = map[string]string{}
	for k, v := range this.GetAll() {
		this.request[k] = v
	}
	for k, v := range this.PostAll() {
		this.request[k] = v
	}
}
func (this *Context) prepareCookie() {
	if nil != this.cookie {
		return
	}
	this.cookie = map[string]string{}
	for _, v := range this.r.Cookies() {
		this.cookie[v.Name] = v.Value
	}
}
func (this *Context) prepareServer() {
	if nil != this.server {
		return
	}
	this.server = map[string]string{}
	for k, v := range this.r.Header {
		this.server[k] = v[0]
	}
	this.server["method"] = this.r.Method
	this.server["path"] = this.r.URL.Path
}

func (this *Context) Get(name string) *_parameter.Parameter {
	this.prepareGet()
	return _parameter.New(name, this.get[name])
}
func (this *Context) GetAll() map[string]string {
	this.prepareGet()
	return this.get
}
func (this *Context) Post(name string) *_parameter.Parameter {
	this.preparePost()
	return _parameter.New(name, this.post[name])
}
func (this *Context) PostAll() map[string]string {
	this.preparePost()
	return this.post
}
func (this *Context) Request(name string) *_parameter.Parameter {
	this.prepareRequest()
	return _parameter.New(name, this.request[name])
}
func (this *Context) RequestAll() map[string]string {
	this.prepareRequest()
	return this.request
}
func (this *Context) Cookie(name string) *_parameter.Parameter {
	this.prepareCookie()
	return _parameter.New(name, this.cookie[name])
}
func (this *Context) CookieAll() map[string]string {
	this.prepareCookie()
	return this.cookie
}
func (this *Context) Server(name string) *_parameter.Parameter {
	this.prepareServer()
	return _parameter.New(name, this.server[name])
}
func (this *Context) ServerAll() map[string]string {
	this.prepareServer()
	return this.server
}
func (this *Context) File(name string) *multipart.FileHeader {
	if f, ok := this.r.MultipartForm.File[name]; ok && len(f) > 0 {
		return f[0]
	}
	return nil
}

func (this *Context) Json(value interface{}) {
	res := _response.New()
	if _is.NotEmpty(value) {
		res.Data = value
	}
	_render.New(this.w, this.r).Json(res)
}
func (this *Context) Text(value interface{}) {
	_render.New(this.w, this.r).Text(value)
}
