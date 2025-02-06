package _base64Format

import (
	"github.com/junyang7/go-common/_assert"
	"testing"
)

func TestEncode(t *testing.T) {
	{
		var expect string = "aGVsbG8gd29ybGQh"
		var give string = "hello world!"
		get := Encode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "5oKo5aW977yM5Lit5Zu977yB"
		var give string = "您好，中国！"
		get := Encode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "aGVsbG8g5Lit5Zu9"
		var give string = "hello 中国"
		get := Encode(give)
		_assert.Equal(t, expect, get)
	}
}
func TestDecode(t *testing.T) {
	{
		var expect string = "hello world!"
		var give string = "aGVsbG8gd29ybGQh"
		get := Decode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "您好，中国！"
		var give string = "5oKo5aW977yM5Lit5Zu977yB"
		get := Decode(give)
		_assert.Equal(t, expect, get)
	}
	{
		var expect string = "hello 中国"
		var give string = "aGVsbG8g5Lit5Zu9"
		get := Decode(give)
		_assert.Equal(t, expect, get)
	}
}
