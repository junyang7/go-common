package _client

import (
	"bytes"
	"context"
	"crypto/tls"
	"git.ziji.fun/junyang/go-common/_as"
	pb2 "git.ziji.fun/junyang/go-common/_client/pb"
	"git.ziji.fun/junyang/go-common/_is"
	"git.ziji.fun/junyang/go-common/_json"
	"git.ziji.fun/junyang/go-common/_list"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Rpc(addr string, header map[string]string, body map[string]string) []byte {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if nil != err {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := pb2.NewServiceClient(conn).Call(ctx, &pb2.Request{
		Header: header,
		Body:   body,
	})
	if nil != err {
		panic(err)
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
						panic(err)
					}
				}
				for name, path := range this.file {
					f, err := os.Open(path)
					if nil != err {
						panic(err)
					}
					defer f.Close()
					fw, err := bw.CreateFormFile(name, path)
					if nil != err {
						panic(err)
					}
					if _, err = io.Copy(fw, f); nil != err {
						panic(err)
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
		panic(err)
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
		panic(err)
	}
	if !_list.In(res.StatusCode, this.httpStatusCode) {
		panic(res.StatusCode)
	}
	if this.needBody {
		defer res.Body.Close()
		b, err := io.ReadAll(res.Body)
		if nil != err {
			panic(err)
		}
		return string(b)
	}
	return ""
}
