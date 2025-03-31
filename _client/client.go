package _client

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/junyang7/go-common/_as"
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_is"
	"github.com/junyang7/go-common/_json"
	"github.com/junyang7/go-common/_list"
	"github.com/junyang7/go-common/_pb"
	"google.golang.org/grpc"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Rpc(addr string, header map[string]string, body []byte) []byte {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := _pb.NewServiceClient(conn).Call(ctx, &_pb.Request{
		Header: header,
		Body:   body,
	})
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return r.GetResponse()
}

const (
	XWwwFormUrlencoded = "application/x-www-form-urlencoded"
	FormData           = "multipart/form-data"
	Json               = "application/json"
	GET                = "GET"
	POST               = "POST"
	PUT                = "PUT"
	DELETE             = "DELETE"
)

type _http struct {
	contentType    string
	method         string
	url            string
	data           map[string]interface{}
	header         map[string]string
	cookie         []*http.Cookie
	file           map[string]string
	httpStatusCode []int
	needBody       bool
}

func Http() *_http {
	return &_http{
		contentType:    XWwwFormUrlencoded,
		method:         GET,
		header:         map[string]string{},
		cookie:         []*http.Cookie{},
		data:           map[string]interface{}{},
		file:           map[string]string{},
		httpStatusCode: []int{200},
		needBody:       true,
	}
}
func (this *_http) Method(method string) *_http {
	this.method = strings.ToUpper(method)
	return this
}
func (this *_http) Url(url string) *_http {
	this.url = url
	return this
}
func (this *_http) ContentType(contentType string) *_http {
	this.contentType = contentType
	return this
}
func (this *_http) Header(name string, value string) *_http {
	this.header[strings.ToLower(name)] = value
	return this
}
func (this *_http) Cookie(cookie *http.Cookie) *_http {
	this.cookie = append(this.cookie, cookie)
	return this
}
func (this *_http) Data(data map[string]interface{}) *_http {
	this.data = data
	return this
}
func (this *_http) File(name string, path string) *_http {
	this.file[name] = path
	return this
}
func (this *_http) HttpStatusCode(httpStatusCode []int) *_http {
	this.httpStatusCode = httpStatusCode
	return this
}
func (this *_http) NeedBody(needBody bool) *_http {
	this.needBody = needBody
	return this
}
func (this *_http) Do() string {
	_url := this.url
	_body := []byte{}
	switch this.method {
	case "GET":
		this.contentType = XWwwFormUrlencoded
		if !_is.Empty(this.data) {
			parameter := url.Values{}
			for k, v := range this.data {
				parameter.Set(k, _as.String(v))
			}
			_url += "?" + parameter.Encode()
		}
		break
	default:
		switch this.contentType {
		case XWwwFormUrlencoded:
			if !_is.Empty(this.data) {
				parameter := url.Values{}
				for k, v := range this.data {
					parameter.Set(k, _as.String(v))
				}
				_body = []byte(parameter.Encode())
			}
			break
		case FormData:
			if !_is.Empty(this.data) || _is.Empty(this.file) {
				bf := &bytes.Buffer{}
				bw := multipart.NewWriter(bf)
				for k, v := range this.data {
					if err := bw.WriteField(k, _as.String(v)); nil != err {
						_interceptor.Insure(false).Message(err).Do()
					}
				}
				for name, path := range this.file {
					f, err := os.Open(path)
					if nil != err {
						_interceptor.Insure(false).Message(err).Do()
					}
					defer f.Close()
					fw, err := bw.CreateFormFile(name, path)
					if nil != err {
						_interceptor.Insure(false).Message(err).Do()
					}
					if _, err = io.Copy(fw, f); nil != err {
						_interceptor.Insure(false).Message(err).Do()
					}
				}
				_ = bw.Close()
				this.contentType = bw.FormDataContentType()
				_body = bf.Bytes()
			}
			break
		case Json:
			if !_is.Empty(this.data) {
				_body = _json.Encode(this.data)
			}
		}
		break
	}
	req, err := http.NewRequest(this.method, _url, bytes.NewReader(_body))
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	req.Header.Set("content-type", this.contentType)
	for k, v := range this.header {
		req.Header.Set(k, v)
	}
	for _, cookie := range this.cookie {
		req.AddCookie(cookie)
	}
	http.DefaultClient.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	res, err := http.DefaultClient.Do(req)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	if !_list.In(res.StatusCode, this.httpStatusCode) {
		_interceptor.Insure(false).Message(res.StatusCode).Do()
	}
	if this.needBody {
		defer res.Body.Close()
		b, err := io.ReadAll(res.Body)
		if nil != err {
			_interceptor.Insure(false).Message(err).Do()
		}
		return string(b)
	}
	return ""
}
