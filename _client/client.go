package _client

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_list"
	"github.com/junyang7/go-common/_pb"
	"github.com/junyang7/go-common/_uri"
	"google.golang.org/grpc"
	"io"
	"mime/multipart"
	netHttp "net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type rpc struct {
	addr    string
	header  map[string]string
	body    []byte
	timeout time.Duration
}

func NewRpc() *rpc {
	return &rpc{
		timeout: 30 * time.Second,
		header:  map[string]string{},
	}
}
func (this *rpc) Addr(addr string) *rpc {
	this.addr = addr
	return this
}
func (this *rpc) Header(k string, v string) *rpc {
	this.header[strings.ToLower(k)] = v
	return this
}
func (this *rpc) Body(body []byte) *rpc {
	this.body = body
	return this
}
func (this *rpc) Timeout(timeout time.Duration) *rpc {
	this.timeout = timeout
	return this
}
func (this *rpc) Do() []byte {
	conn, err := grpc.Dial(this.addr, grpc.WithInsecure())
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	ctx, cancel := context.WithTimeout(context.Background(), this.timeout)
	defer cancel()
	r, err := _pb.NewServiceClient(conn).Call(ctx, &_pb.Request{
		Header: this.header,
		Body:   this.body,
	})
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return r.GetResponse()
}
func Rpc(addr string, header map[string]string, body []byte, timeout time.Duration) []byte {
	c := NewRpc()
	c.Addr(addr)
	for k, v := range header {
		c.Header(k, v)
	}
	c.Body(body)
	c.Timeout(timeout)
	return c.Do()
}

const (
	XWwwFormUrlencoded = "application/x-www-form-urlencoded"
	FormData           = "multipart/form-data"
	Json               = "application/json"
)

type http struct {
	contentType string
	timeout     time.Duration
	statusCode  []int
	method      string
	url         string
	cookie      []*netHttp.Cookie
	header      map[string]string
	data        map[string]string
	file        map[string]string
	body        []byte
	v           interface{}
}

func NewHttp() *http {
	return &http{
		contentType: XWwwFormUrlencoded,
		timeout:     30 * time.Second,
		statusCode:  []int{200},
		method:      netHttp.MethodGet,
		url:         "",
		cookie:      []*netHttp.Cookie{},
		header:      map[string]string{},
		data:        map[string]string{},
		file:        map[string]string{},
	}
}
func (this *http) ContentType(contentType string) *http {
	this.contentType = strings.TrimSpace(contentType)
	return this
}
func (this *http) Timeout(timeout time.Duration) *http {
	this.timeout = timeout
	return this
}
func (this *http) StatusCode(statusCode ...int) *http {
	this.statusCode = statusCode
	return this
}
func (this *http) Method(method string) *http {
	this.method = strings.ToUpper(method)
	_interceptor.
		Insure(_list.In(this.method, []string{netHttp.MethodGet, netHttp.MethodPost})).
		Data(map[string]interface{}{"method": this.method}).
		Message("不支持的请求方法").
		Do()
	return this
}
func (this *http) Url(url string) *http {
	this.url = strings.TrimSpace(url)
	return this
}
func (this *http) Cookie(cookie *netHttp.Cookie) *http {
	this.cookie = append(this.cookie, cookie)
	return this
}
func (this *http) Header(name string, value string) *http {
	this.header[strings.ToLower(strings.TrimSpace(name))] = value
	return this
}
func (this *http) Data(data map[string]string) *http {
	this.data = data
	return this
}
func (this *http) File(name string, path string) *http {
	this.file[name] = path
	return this
}
func (this *http) Bind(v interface{}) *http {
	this.v = v
	return this
}
func (this *http) Do() []byte {
	_interceptor.Insure(this.url != "").Message("url不能为空").Do()
	url := this.url
	body := []byte{}
	if this.contentType != "" {
		this.header["content-type"] = this.contentType
	}
	if this.method == netHttp.MethodGet {
		if len(this.data) > 0 {
			url = url + "?" + _uri.Build(this.data)
		}
	} else if this.method == netHttp.MethodPost {
		if strings.HasPrefix(this.header["content-type"], FormData) {
			bf := &bytes.Buffer{}
			bw := multipart.NewWriter(bf)
			for k, v := range this.data {
				_ = bw.WriteField(k, v)
			}
			for k, v := range this.file {
				func() {
					f, err := os.Open(v)
					if err != nil {
						_interceptor.Insure(false).Message(err).Do()
					}
					defer f.Close()
					fw, _ := bw.CreateFormFile(k, filepath.Base(v))
					_, _ = io.Copy(fw, f)
				}()
			}
			_ = bw.Close()
			body = bf.Bytes()
			this.header["content-type"] = bw.FormDataContentType()
		} else if strings.HasPrefix(this.header["content-type"], Json) {
			if len(this.body) > 0 {
				body = this.body
			} else if len(this.data) > 0 {
				body = _json.Encode(this.data)
			}
		} else {
			if len(this.data) > 0 {
				body = []byte(_uri.Build(this.data))
			}
		}
	}
	req, err := netHttp.NewRequest(this.method, url, bytes.NewReader(body))
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	for _, cookie := range this.cookie {
		req.AddCookie(cookie)
	}
	for k, v := range this.header {
		req.Header.Set(k, v)
	}
	client := &netHttp.Client{
		Timeout: this.timeout,
		Transport: &netHttp.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	res, err := client.Do(req)
	_interceptor.
		Insure(nil == err).
		Message(err).
		Do()
	_interceptor.
		Insure(_list.In(res.StatusCode, this.statusCode)).
		Data(map[string]interface{}{"state_code": res.StatusCode}).
		Message(res.StatusCode).
		Do()
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	if strings.HasPrefix(strings.ToLower(res.Header.Get("content-type")), Json) {
		_json.Decode(b, this.v)
	}
	return b
}
