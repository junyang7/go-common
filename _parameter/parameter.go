package _parameter

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_validator"
	"mime/multipart"
)

type Parameter struct {
	name  string      // 参数名
	value interface{} // 参数值
}

func New(name string, value interface{}) *Parameter {
	return &Parameter{
		name:  name,
		value: value,
	}
}
func (p *Parameter) Default(value interface{}) *Parameter {
	if p.value == nil {
		p.value = value
	}
	return p
}
func (p *Parameter) Required() *Parameter {
	if p.value == nil {
		_interceptor.Insure(false).
			Message(p.name + " is required").
			Do()
	}
	return p
}
func (p *Parameter) Value() interface{} {
	return p.value
}
func (p *Parameter) Int64() *_validator.Int64 {
	return _validator.NewInt64(p.name, p.value)
}
func (p *Parameter) String() *_validator.String {
	return _validator.NewString(p.name, p.value)
}
func (p *Parameter) Bool() *_validator.Bool {
	return _validator.NewBool(p.name, p.value)
}
func (p *Parameter) Float64() *_validator.Float64 {
	return _validator.NewFloat64(p.name, p.value)
}
func (p *Parameter) File() *multipart.FileHeader {
	if files, ok := p.value.([]*multipart.FileHeader); ok {
		if len(files) > 0 {
			return files[0]
		}
	}
	return nil
}
func (p *Parameter) FileList() []*multipart.FileHeader {
	if files, ok := p.value.([]*multipart.FileHeader); ok {
		return files
	}
	return nil
}
