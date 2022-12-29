package _http

import (
	"bytes"
	"crypto/tls"
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_codeMessage"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_is"
	"github.com/junyang7/go-common/src/_json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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
	contentType string
	method      string
	url         string
	data        map[string]interface{}
	header      map[string]string
	cookie      []*http.Cookie
	file        map[string]string
}

func New() *_http {
	return &_http{
		contentType: XWwwFormUrlencoded,
		method:      GET,
		header:      map[string]string{},
		cookie:      []*http.Cookie{},
		data:        map[string]interface{}{},
		file:        map[string]string{},
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
func (this *_http) Do() string {

	data := map[string]interface{}{"method": this.method, "url": this.url, "content-type": this.contentType, "header": this.header, "cookie": this.cookie, "data": this.data, "file": this.file}
	_url := this.url
	var _body []byte

	switch this.method {
	case "GET":
		this.contentType = XWwwFormUrlencoded
		if _is.NotEmpty(this.data) {
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
			if _is.NotEmpty(this.data) {
				parameter := url.Values{}
				for k, v := range this.data {
					parameter.Set(k, _as.String(v))
				}
				_body = []byte(parameter.Encode())
			}
			break
		case FormData:
			if _is.NotEmpty(this.data) || _is.Empty(this.file) {
				bf := &bytes.Buffer{}
				bw := multipart.NewWriter(bf)
				for k, v := range this.data {
					err := bw.WriteField(k, _as.String(v))
					_interceptor.Insure(nil == err).
						CodeMessage(_codeMessage.ErrMultipartNewWriterWriteField).
						Data(data).
						Do()
				}
				for name, path := range this.file {
					f, err := os.Open(path)
					_interceptor.Insure(nil == err).
						CodeMessage(_codeMessage.ErrOsOpen).
						Data(map[string]interface{}{"path": path}).
						Do()
					defer f.Close()
					fw, err := bw.CreateFormFile(name, path)
					_interceptor.Insure(nil == err).
						CodeMessage(_codeMessage.ErrMultipartNewWriterCreateFormFile).
						Data(map[string]interface{}{"name": name, "path": path}).
						Do()
					_, err = io.Copy(fw, f)
					_interceptor.Insure(nil == err).
						CodeMessage(_codeMessage.ErrIoCopy).
						Data(map[string]interface{}{"name": name, "path": path}).
						Do()
				}
				_ = bw.Close()
				this.contentType = bw.FormDataContentType()
				_body = bf.Bytes()
			}
			break
		case Json:
			if _is.NotEmpty(this.data) {
				_body = _json.Encode(this.data)
			}
		}
		break
	}

	req, err := http.NewRequest(this.method, _url, bytes.NewReader(_body))
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHttpNewRequest).
		Data(data).
		Do()

	req.Header.Set("content-type", this.contentType)
	for k, v := range this.header {
		req.Header.Set(k, v)
	}

	for _, cookie := range this.cookie {
		req.AddCookie(cookie)
	}

	http.DefaultClient.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	res, err := http.DefaultClient.Do(req)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrHttpDefaultClientDo).
		Data(data).
		Do()

	defer res.Body.Close()

	_interceptor.Insure(http.StatusOK == res.StatusCode).
		CodeMessage(_codeMessage.ErrHttpStateCodeNot200).
		Data(data).
		Do()

	b, err := ioutil.ReadAll(res.Body)
	_interceptor.Insure(nil == err).
		CodeMessage(_codeMessage.ErrIoUtilReadAll).
		Data(data).
		Do()

	return string(b)

}
