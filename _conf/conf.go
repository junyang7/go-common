package _conf

import "github.com/junyang7/go-common/_parameter"

type Conf interface {
	Byte(byte []byte) Conf
	Text(text string) Conf
	File(path string) Conf
	Get(path string) *_parameter.Parameter
}
