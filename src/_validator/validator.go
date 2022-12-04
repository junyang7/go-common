package _validator

import "github.com/junyang7/go-common/src/_as"

type validator struct {
	name  string
	value string
}

func New(name string, value string) *validator {
	return &validator{
		name:  name,
		value: value,
	}
}
func (this *validator) Int() int {
	return _as.Int(this.value)
}
func (this *validator) String() string {
	return _as.String(this.value)
}
func (this *validator) Bool() bool {
	return _as.Bool(this.value)
}
func (this *validator) Float64() float64 {
	return _as.Float64(this.value)
}
