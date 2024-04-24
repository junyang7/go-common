package _parameter

import (
	"git.ziji.fun/junyang/go-common/_validator"
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
