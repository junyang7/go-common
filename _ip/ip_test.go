package _ip

import (
	"github.com/junyang7/go-common/_assert"
	"regexp"
	"testing"
)

func TestGetByLocal(t *testing.T) {
	{
		ip := GetByLocal()
		matchedList := regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`).FindStringSubmatch(ip)
		_assert.Equal(t, 1, len(matchedList))
	}
}
func TestGetListByLocal(t *testing.T) {
	{
		ipList := GetListByLocal()
		_assert.True(t, len(ipList) > 0)
	}
}
