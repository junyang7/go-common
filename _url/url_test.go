package _url

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestParse(t *testing.T) {
	{
		url := "https://www.example.com/path?query=1"
		parsed := Parse(url)
		_assert.Equal(t, "https", parsed.Scheme)
		_assert.Equal(t, "www.example.com", parsed.Host)
		_assert.Equal(t, "/path", parsed.Path)
		_assert.Equal(t, "query=1", parsed.RawQuery)
		_assert.Equal(t, "https://www.example.com/path?query=1", parsed.String())
	}
	{
		url := "https://www.example.com"
		parsed := Parse(url)
		_assert.Equal(t, "https", parsed.Scheme)
		_assert.Equal(t, "www.example.com", parsed.Host)
		_assert.Equal(t, "", parsed.Path)
		_assert.Equal(t, "", parsed.RawQuery)
		_assert.Equal(t, "https://www.example.com", parsed.String())
	}
	{
		url := "ftp://ftp.example.com/file"
		parsed := Parse(url)
		_assert.Equal(t, "ftp", parsed.Scheme)
		_assert.Equal(t, "ftp.example.com", parsed.Host)
		_assert.Equal(t, "/file", parsed.Path)
		_assert.Equal(t, "", parsed.RawQuery)
		_assert.Equal(t, "ftp://ftp.example.com/file", parsed.String())
	}
}
func TestGetOrigin(t *testing.T) {
	{
		url := "https://www.example.com/path?query=1"
		expected := "https://www.example.com"
		get := GetOrigin(url)
		_assert.Equal(t, expected, get)
	}
	{
		url := "https://www.example.com"
		expected := "https://www.example.com"
		get := GetOrigin(url)
		_assert.Equal(t, expected, get)
	}
	{
		url := "http://www.example.com"
		expected := "http://www.example.com"
		get := GetOrigin(url)
		_assert.Equal(t, expected, get)
	}
	{
		url := "ftp://ftp.example.com/file"
		expected := "ftp://ftp.example.com"
		get := GetOrigin(url)
		_assert.Equal(t, expected, get)
	}
}
