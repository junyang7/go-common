package _url

import (
	"github.com/junyang7/go-common/_interceptor"
	url_ "net/url"
)

func Parse(url string) *url_.URL {
	f, err := url_.Parse(url)
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	return f
}
func GetOrigin(url string) string {
	f := Parse(url)
	return f.Scheme + "://" + f.Host
}
