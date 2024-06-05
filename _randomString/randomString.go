package _randomString

import (
	"math/rand"
	"time"
)

//public static function get($size = 32, $string = "0123456789QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm~!@#$%^&*()_+[]\;,./{}:<>?|")
//    {
//
//        $res = "";
//        $index_max = strlen($string) - 1;
//        for ($i = 0; $i < $size; $i++) {
//            $res .= $string[rand(0, $index_max)];
//        }
//        return $res;
//
//    }

type randomString struct {
	size int
	char string
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
func (this *randomString) Get() string {
	rand.Seed(time.Now().UnixNano())
	res := ``
	for i, j := 0, len(this.char); i < this.size; i++ {
		k := rand.Intn(j)
		res += this.char[k : k+1]
	}
	return res
}
