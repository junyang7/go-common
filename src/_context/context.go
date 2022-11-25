package _context

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strings"
)

type Context struct {
	W       http.ResponseWriter
	R       *http.Request
	get     map[string]string
	post    map[string]string
	request map[string]string
	cookie  map[string]string
	server  map[string]string
}

func (this *Context) prepareGet() {
	if nil != this.get {
		return
	}
	this.get = map[string]string{}
	for k, v := range this.R.URL.Query() {
		this.get[k] = v[0]
	}
}
func (this *Context) preparePost() {
	if nil != this.post {
		return
	}
	this.post = map[string]string{}
	if err := this.R.ParseForm(); nil != err {
		panic(err)
	}
	contentType := strings.ToLower(this.R.Header.Get("content-type"))
	if -1 != strings.Index(contentType, "application/x-www-form-urlencoded") {
		for k, v := range this.R.PostForm {
			this.post[k] = v[0]
		}
	}
	if -1 != strings.Index(contentType, "multipart/form-data") {
		if err := this.R.ParseMultipartForm(32 << 20); nil != err {
			panic(err)
		}
		for k, v := range this.R.PostForm {
			this.post[k] = v[0]
		}
	}
	if -1 != strings.Index(contentType, "application/json") {
		if err := json.NewDecoder(this.R.Body).Decode(&this.post); nil != err {
			panic(err)
		}
		if err := this.R.Body.Close(); nil != err {
			panic(err)
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
	for _, v := range this.R.Cookies() {
		this.cookie[v.Name] = v.Value
	}
}
func (this *Context) prepareServer() {
	if nil != this.server {
		return
	}
	this.server = map[string]string{}
	for k, v := range this.R.Header {
		this.server[k] = v[0]
	}
	this.server["method"] = this.R.Method
	this.server["path"] = this.R.URL.Path
}

func (this *Context) Get(name string) string {
	this.prepareGet()
	if value, ok := this.get[name]; ok {
		return value
	}
	return ""
}
func (this *Context) GetAll() map[string]string {
	this.prepareGet()
	return this.get
}
func (this *Context) Post(name string) string {
	this.preparePost()
	if value, ok := this.post[name]; ok {
		return value
	}
	return ""
}
func (this *Context) PostAll() map[string]string {
	this.preparePost()
	return this.post
}
func (this *Context) Request(name string) string {
	this.prepareRequest()
	if value, ok := this.request[name]; ok {
		return value
	}
	return ""
}
func (this *Context) RequestAll() map[string]string {
	this.prepareRequest()
	return this.request
}
func (this *Context) Cookie(name string) string {
	this.prepareCookie()
	if value, ok := this.cookie[name]; ok {
		return value
	}
	return ""
}
func (this *Context) CookieAll() map[string]string {
	this.prepareCookie()
	return this.cookie
}
func (this *Context) Server(name string) string {
	this.prepareServer()
	if value, ok := this.server[name]; ok {
		return value
	}
	return ""
}
func (this *Context) ServerAll() map[string]string {
	this.prepareServer()
	return this.server
}
func (this *Context) File(name string) *multipart.FileHeader {
	if f, ok := this.R.MultipartForm.File[name]; ok && len(f) > 0 {
		return f[0]
	}
	return nil
}

func (this *Context) Json(data interface{}) {
	this.W.Header().Set("content-type", "application/json")
	e := json.NewEncoder(this.W)
	_ = e.Encode(data)
}
func (this *Context) String(data interface{}) {
	this.W.Header().Set("content-type", "application/text")
	e := json.NewEncoder(this.W)
	_ = e.Encode(data)
}
