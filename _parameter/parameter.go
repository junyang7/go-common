package _parameter

import (
	"github.com/junyang7/go-common/_validator"
	"os"
	"strings"
)

type Parameter struct {
	name  string
	value interface{}
}

func New(name string, value interface{}) *Parameter {
	return &Parameter{
		name:  name,
		value: value,
	}
}
func (this *Parameter) Int() *_validator.Int {
	return _validator.NewInt(this.name, this.value)
}
func (this *Parameter) String() *_validator.String {
	return _validator.NewString(this.name, this.value)
}
func (this *Parameter) Bool() *_validator.Bool {
	return _validator.NewBool(this.name, this.value)
}
func (this *Parameter) Float64() *_validator.Float64 {
	return _validator.NewFloat64(this.name, this.value)
}
func ParseByTerminal() map[string]string {
	res := map[string]string{}
	for k, v := range os.Args {
		if 0 == k {
			res["entrypoint"] = v
			continue
		}
		partList := strings.Split(v, "=")
		if 2 == len(partList) {
			res[strings.Trim(partList[0], "-")] = partList[1]
		}
	}
	return res
}
