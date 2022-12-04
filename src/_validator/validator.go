package _validator

import (
	"github.com/junyang7/go-common/src/_as"
	"github.com/junyang7/go-common/src/_is"
)

type Validator struct {
	name  string
	value interface{}
}

func New(name string, value interface{}) *Validator {
	return &Validator{
		name:  name,
		value: value,
	}
}
func (this *Validator) Default(value interface{}) *Validator {
	if _is.Empty(this.value) {
		this.value = value
	}
	return this
}
func (this *Validator) Int() int {
	return _as.Int(this.value)
}
func (this *Validator) String() string {
	return _as.String(this.value)
}
func (this *Validator) Bool() bool {
	return _as.Bool(this.value)
}
func (this *Validator) Float64() float64 {
	return _as.Float64(this.value)
}
