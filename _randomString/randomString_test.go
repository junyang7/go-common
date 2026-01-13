package _randomString

import (
	"github.com/junyang7/go-common/_assert"
	"github.com/junyang7/go-common/_is"
	"strings"
	"testing"
)

func TestGetByNumber(t *testing.T) {
	{
		result := GetByNumber(10000)
		_assert.Equal(t, len(result), 10000)
		_assert.True(t, _is.Numeric(result))
	}
}
func TestGetByAlpha(t *testing.T) {
	{
		result := GetByAlpha(10000)
		_assert.Equal(t, len(result), 10000)
		_assert.True(t, _is.Alpha(result))
	}
}
func TestGetByAlphaLower(t *testing.T) {
	{
		result := GetByAlphaLower(10000)
		_assert.Equal(t, len(result), 10000)
		_assert.True(t, _is.AlphaLower(result))
	}
}
func TestGetByAlphaUpper(t *testing.T) {
	{
		result := GetByAlphaUpper(10000)
		_assert.Equal(t, len(result), 10000)
		_assert.True(t, _is.AlphaUpper(result))
	}
}
func TestGetByNumberAndAlpha(t *testing.T) {
	{
		result := GetByNumberAndAlpha(10000)
		_assert.Equal(t, len(result), 10000)
		_assert.True(t, strings.ContainsAny(result, "0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"))
	}
}
func TestGetByNumberAndAlphaAndSpecialChar(t *testing.T) {
	{
		result := GetByNumberAndAlphaAndSpecialChar(10000)
		_assert.Equal(t, len(result), 10000)
		_assert.True(t, len(result) > 0)
		_assert.True(t, strings.ContainsAny(result, `0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$%^&*()_+[]\;,./{}:<>?|`))
	}
}
func TestGet(t *testing.T) {
	{
		random := New().Size(10000).Char("abcdefghigklmn1234567890").Filter("a").Filter("b").FilterList([]string{"c", "d"})
		result := random.Get()
		_assert.Equal(t, len(result), 10000)
		_assert.False(t, strings.Contains(result, "a"))
		_assert.False(t, strings.Contains(result, "b"))
	}
}
