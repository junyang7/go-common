package _randomString

import (
	"math/rand"
	"strings"
	"time"
)

func Get(size int, char string) string {
	rand.Seed(time.Now().UnixNano())
	res := ``
	for i, j := 0, len(char); i < size; i++ {
		k := rand.Intn(j)
		res += char[k : k+1]
	}
	return res
}
func GetByNumber(size int) string {
	return Get(size, `0123456789`)
}
func GetByAlpha(size int) string {
	return Get(size, `QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm`)
}
func GetByAlphaLower(size int) string {
	return Get(size, `qwertyuiopasdfghjklzxcvbnm`)
}
func GetByAlphaUpper(size int) string {
	return Get(size, `QWERTYUIOPASDFGHJKLZXCVBNM`)
}
func GetByNumberAndAlpha(size int) string {
	return Get(size, `0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm`)
}
func GetByNumberAndAlphaAndSpecialChar(size int) string {
	return Get(size, `0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$%^&*()_+[]\;,./{}:<>?|`)
}

type randomString struct {
	size       int
	char       string
	filterList []string
}

func New() *randomString {
	return &randomString{
		size: 16,
		char: `0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$%^&*()_+[]\;,./{}:<>?|`,
	}
}
func (this *randomString) Size(size int) *randomString {
	if size > 0 {
		this.size = size
	}
	return this
}
func (this *randomString) Char(char string) *randomString {
	if len(char) > 0 {
		this.char = char
	}
	return this
}
func (this *randomString) FilterList(filterList []string) *randomString {
	if len(filterList) > 0 {
		this.filterList = append(this.filterList, filterList...)
	}
	return this
}
func (this *randomString) Filter(filter string) *randomString {
	if "" != filter {
		this.filterList = append(this.filterList, filter)
	}
	return this
}
func (this *randomString) Get() string {
	filteredChar := this.char
	for _, filter := range this.filterList {
		filteredChar = strings.ReplaceAll(filteredChar, filter, "")
	}
	rand.Seed(time.Now().UnixNano())
	res := ``
	for i, j := 0, len(filteredChar); i < this.size; i++ {
		k := rand.Intn(j)
		res += filteredChar[k : k+1]
	}
	return res
}
