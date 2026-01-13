package _uri

import "net/url"

func Build(parameter map[string]string) string {
	uri := url.Values{}
	for k, v := range parameter {
		uri.Set(k, v)
	}
	return uri.Encode()
}
