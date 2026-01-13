package _us

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestEncodeAndDecode(t *testing.T) {
	{
		conf := &Conf{
			K32:      "0123456789abcdef0123456789abcdef",
			I16:      "0123456789abcdef",
			Expires:  3600,
			Path:     "/",
			Domain:   "example.com",
			Secure:   "true",
			HttpOnly: "true",
		}
		data := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		encoded := Encode(data, conf)
		decoded := Decode(encoded[:len(encoded)-32], conf)
		for key, val := range data {
			_assert.In(t, key, []string{"_r", "_t", "_e", "key1", "key2"})
			_assert.Equal(t, decoded[key], val)
		}
	}
}
