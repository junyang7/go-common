package _ip

import (
	"git.ziji.fun/junyang/go-common/_assert"
	"regexp"
	"testing"
)

func TestGetByLocal(t *testing.T) {
	{
		ip := GetByLocal()
		matchedList := regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+$`).FindStringSubmatch(ip)
		var expect int = 1
		get := len(matchedList)
		_assert.Equal(t, expect, get)
	}
}
