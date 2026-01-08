package _parameter

import (
	"github.com/junyang7/go-common/_interceptor"
	"github.com/junyang7/go-common/_validator"
	"mime/multipart"
	"os"
	"strings"
)

type Parameter struct {
	name  string
	value interface{}
}

// New 创建参数
func New(name string, value interface{}) *Parameter {
	return &Parameter{
		name:  name,
		value: value,
	}
}

// Default 设置默认值（仅当 value 为 nil 时）
func (p *Parameter) Default(value interface{}) *Parameter {
	if p.value == nil {
		p.value = value
	}
	return p
}

// Required 必填验证
func (p *Parameter) Required() *Parameter {
	if p.value == nil {
		_interceptor.Insure(false).
			Message(p.name + " is required").
			Do()
	}
	return p
}

// Value 获取原始值
func (p *Parameter) Value() interface{} {
	return p.value
}

// ============================================================
// 类型转换（统一使用 Validator）
// ============================================================

// Int64 转换为 Int64 Validator
func (p *Parameter) Int64() *_validator.Int64 {
	return _validator.NewInt64(p.name, p.value)
}

// String 转换为 String Validator
func (p *Parameter) String() *_validator.String {
	return _validator.NewString(p.name, p.value)
}

// Bool 转换为 Bool Validator
func (p *Parameter) Bool() *_validator.Bool {
	return _validator.NewBool(p.name, p.value)
}

// Float64 转换为 Float64 Validator
func (p *Parameter) Float64() *_validator.Float64 {
	return _validator.NewFloat64(p.name, p.value)
}

// ============================================================
// 文件支持
// ============================================================

// File 获取第一个文件
func (p *Parameter) File() *multipart.FileHeader {
	if files, ok := p.value.([]*multipart.FileHeader); ok {
		if len(files) > 0 {
			return files[0]
		}
	}
	return nil
}

// FileList 获取文件列表
func (p *Parameter) FileList() []*multipart.FileHeader {
	if files, ok := p.value.([]*multipart.FileHeader); ok {
		return files
	}
	return nil
}

// ============================================================
// 辅助方法
// ============================================================

// IsNil 检查值是否为 nil
func (p *Parameter) IsNil() bool {
	return p.value == nil
}

// ============================================================
// 命令行参数解析
// ============================================================

// ParseByTerminal 解析命令行参数
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
